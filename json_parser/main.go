package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TokenType is just an int under the hood — we use "iota" to auto-assign
// incrementing values: EOF=0, LBRACE=1, RBRACE=2, etc.
type TokenType int

const (
	TOKEN_EOF      TokenType = iota // 0 — we've consumed the entire input
	TOKEN_LBRACE                    // 1 — {
	TOKEN_RBRACE                    // 2 — }
	TOKEN_LBRACKET                  // 3 — [
	TOKEN_RBRACKET                  // 4 — ]
	TOKEN_COMMA                     // 5 — ,
	TOKEN_COLON                     // 6 — :
	TOKEN_STRING                    // 7 — "hello"
	TOKEN_NUMBER                    // 8 — 42 or 3.14
	TOKEN_TRUE                      // 9 — true
	TOKEN_FALSE                     // 10 — false
	TOKEN_NULL                      // 11 — null
)

// Token holds the type and the actual text slice from the input.
// e.g. for `"hello"` → {Type: TOKEN_STRING, Value: "hello"}
type Token struct {
	Type  TokenType
	Value string
}

// Lexer owns the input string and a cursor (pos) that advances as we consume chars.
type Lexer struct {
	input string
	pos   int // index of the next character to be read
}

type Parser struct {
	lexer   *Lexer
	current Token // the token we justt consument/ are currently looking at
}

// NewLexer constructs a Lexer positioned at the start of the input.
func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func NewParser(input string) *Parser {
	p := &Parser{lexer: NewLexer(input)}
	p.parserAdvance()
	return p
}

// parserAdvance() moves to the next token - like turning page.
// After this call, p.current holds the token we just moved to
func (p *Parser) parserAdvance() Token {
	prev := p.current
	p.current = p.lexer.GetNextToken()
	return prev // return what we just left behind
}

// consume asserts that the current token is what we expect, then advances
//if the type doesn't match something is malformed - we return an error
//This is the consume(TOKEN_X) from pseudocode

func (p *Parser) consume(expected TokenType) (Token, error) {
	if p.current.Type != expected {
		return Token{}, fmt.Errorf("expected token %d, got %d", expected, p.current.Type)
	}
	return p.parserAdvance(), nil
}

// ParseValue is the heart of the parser — it looks at the current token
// and dispatches to the right handler. This mirrors the pseudo-code's CASE block.
// Return type is `any` because a JSON value can be: string, float64, bool, nil,
// map[string]any (object), or []any (array).
func (p *Parser) ParseValue() (any, error) {
	tok := p.current
	switch tok.Type {
	case TOKEN_STRING:
		p.parserAdvance()
		return tok.Value, nil
	case TOKEN_NUMBER:
		p.parserAdvance()
		//strconv, ParseFloat turns 3.14 to float64
		// JSON has no int/float distinction - all members are float64
		num, err := strconv.ParseFloat(tok.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", tok.Value, err)
		}
		return num, nil

	case TOKEN_TRUE:
		p.parserAdvance()
		return true, nil

	case TOKEN_FALSE:
		p.parserAdvance()
		return false, nil

	case TOKEN_NULL:
		p.parserAdvance()
		return nil, nil

	case TOKEN_LBRACE:
		return p.ParseObject() // delegate

	case TOKEN_LBRACKET:
		return p.ParseArray() // delegate
	default:
		return nil, fmt.Errorf("unexpected token %q at this position", tok.Value)
	}
}
func printValue(v any, depth int) {
	indent := fmt.Sprintf("%*s", depth*2, "") // depth*2 spaces

	switch val := v.(type) {
	case map[string]any:
		fmt.Println(indent + "{")
		for k, v := range val {
			fmt.Printf("%s  %q: ", indent, k)
			printValue(v, depth+1)
		}
		fmt.Println(indent + "}")
	case []any:
		fmt.Println(indent + "[")
		for _, item := range val {
			printValue(item, depth+1)
		}
		fmt.Println(indent + "]")
	default:
		fmt.Printf("%v\n", val)
	}
}

// atEnd returns true when the cursor has moved past the last character.
func (l *Lexer) atEnd() bool {
	return l.pos >= len(l.input)
}

// peek returns the current character WITHOUT advancing pos.
// Used to look at what's coming next before deciding what to do.
func (l *Lexer) peek() byte {
	return l.input[l.pos]
}

// advance returns the current character AND moves pos forward by one.
// Think of it as "consume this character and move on".
func (l *Lexer) advance() byte {
	ch := l.input[l.pos]
	l.pos++
	return ch
}

// skipWhitespace moves pos forward over any spaces, tabs, or newlines.
// JSON doesn't care about whitespace between tokens, so we just skip it.
func (l *Lexer) skipWhitespace() {
	for !l.atEnd() && strings.ContainsRune(" \t\n\r", rune(l.peek())) {
		l.advance()
	}
}

// consumeString handles everything after the opening `"`.
// It reads characters until the closing `"`, building the string value.
func (l *Lexer) consumeString() Token {
	l.advance() // consume the opening `"`

	var sb strings.Builder

	for !l.atEnd() {
		ch := l.advance()

		if ch == '\\' { // escape sequence: \", \\, \n, \t, etc.
			if l.atEnd() {
				break
			}
			escaped := l.advance() // the character after the backslash
			switch escaped {
			case '"':
				sb.WriteByte('"')
			case '\\':
				sb.WriteByte('\\')
			case 'n':
				sb.WriteByte('\n')
			case 't':
				sb.WriteByte('\t')
			case 'r':
				sb.WriteByte('\r')
			default:
				// Unknown escape — write both chars as-is for now
				sb.WriteByte('\\')
				sb.WriteByte(escaped)
			}
			continue
		}

		if ch == '"' { // closing quote — string is done
			break
		}

		sb.WriteByte(ch) // normal character, just accumulate it
	}

	return Token{Type: TOKEN_STRING, Value: sb.String()}
}

// ParseObject  handles everything between { and }
// It builds a Go map[string]any which is natural equivalent of JSON object
func (p *Parser) ParseObject() (map[string]any, error) {
	result := make(map[string]any)
	_, err := p.consume(TOKEN_LBRACE) // consume the opening
	if err != nil {
		return nil, err
	}
	if p.current.Type == TOKEN_RBRACE {
		p.parserAdvance()
		return result, nil
	}

	for p.current.Type != TOKEN_RBRACE {
		// ----parse the key ---
		// JSON keys must be strings. ParseValue is called but we
		// type-assert the result to enforce that constraint
		keyVal, err := p.ParseValue()
		if err != nil {
			return nil, err
		}
		key, ok := keyVal.(string)
		if !ok {
			return nil, fmt.Errorf("object key must be string, got %T", keyVal)
		}
		// --- parse the value (the recursive magic!) ---
		// ParseValue calls itself recursively, so nested objects/arrays just work.
		// e.g. {"a": {"b": 1}}  →  when parsing the value of "a",
		//      ParseValue sees TOKEN_LBRACE and calls ParseObject again.
		_, err = p.consume(TOKEN_COLON)
		if err != nil {
			return nil, err
		}

		value, err := p.ParseValue()
		if err != nil {
			return nil, err
		}
		result[key] = value // store key- > value

		//after each pair, there's either a comma
		//or the closing '}' we're done
		if p.current.Type == TOKEN_COMMA {
			p.parserAdvance()
		}
	}
	_, err = p.consume(TOKEN_RBRACE) // consume the closing `}`
	if err != nil {
		return nil, err
	}

	return result, nil
}

// consumeNumber reads an integer or float.
// JSON numbers can be negative and can have decimals: -3.14
func (l *Lexer) consumeNumber() Token {
	start := l.pos // remember where the number started

	if !l.atEnd() && l.peek() == '-' {
		l.advance() // consume the optional leading minus sign
	}

	// consume all digit characters
	for !l.atEnd() && l.peek() >= '0' && l.peek() <= '9' {
		l.advance()
	}

	// if there's a '.', consume the decimal part too
	if !l.atEnd() && l.peek() == '.' {
		l.advance() // consume '.'
		for !l.atEnd() && l.peek() >= '0' && l.peek() <= '9' {
			l.advance()
		}
	}

	// slice the original input from start→pos gives us the raw number text
	return Token{Type: TOKEN_NUMBER, Value: l.input[start:l.pos]}
}
func (p *Parser) ParseArray() ([]any, error) {
	result := []any{} // empty slice

	_, err := p.consume(TOKEN_LBRACKET) // consume opening `[`
	if err != nil {
		return nil, err
	}

	// Empty array: `[]`
	if p.current.Type == TOKEN_RBRACKET {
		p.parserAdvance()
		return result, nil
	}

	for p.current.Type != TOKEN_RBRACKET {
		// Each element can be any JSON value — string, number, object, another array...
		// Recursion handles arbitrary nesting depth automatically.
		val, err := p.ParseValue()
		if err != nil {
			return nil, err
		}
		result = append(result, val)

		if p.current.Type == TOKEN_COMMA {
			p.parserAdvance() // consume `,` between elements
		}
	}

	_, err = p.consume(TOKEN_RBRACKET) // consume closing `]`
	if err != nil {
		return nil, err
	}

	return result, nil
}

// consumeLiteral tries to match a fixed word like "true", "false", "null".
// If the input at pos starts with the expected word, we advance past it.
func (l *Lexer) consumeLiteral(word string, t TokenType) (Token, bool) {
	end := l.pos + len(word)
	if end <= len(l.input) && l.input[l.pos:end] == word {
		l.pos = end // jump pos forward by the length of the word
		return Token{Type: t, Value: word}, true
	}
	return Token{}, false // doesn't match
}

// GetNextToken is the main entry point — the pseudo-code's core function.
// Each call returns one token and advances pos past it.
func (l *Lexer) GetNextToken() Token {
	l.skipWhitespace() // step 1: skip any leading whitespace

	if l.atEnd() { // nothing left to read
		return Token{Type: TOKEN_EOF}
	}

	ch := l.peek() // look at current char WITHOUT consuming it yet

	// Single-character tokens: just advance and return the matching type.
	switch ch {
	case '{':
		l.advance()
		return Token{Type: TOKEN_LBRACE, Value: "{"}
	case '}':
		l.advance()
		return Token{Type: TOKEN_RBRACE, Value: "}"}
	case '[':
		l.advance()
		return Token{Type: TOKEN_LBRACKET, Value: "["}
	case ']':
		l.advance()
		return Token{Type: TOKEN_RBRACKET, Value: "]"}
	case ',':
		l.advance()
		return Token{Type: TOKEN_COMMA, Value: ","}
	case ':':
		l.advance()
		return Token{Type: TOKEN_COLON, Value: ":"}
	case '"':
		return l.consumeString() // multi-char: delegate to helper
	case 't':
		if tok, ok := l.consumeLiteral("true", TOKEN_TRUE); ok {
			return tok
		}
	case 'f':
		if tok, ok := l.consumeLiteral("false", TOKEN_FALSE); ok {
			return tok
		}
	case 'n':
		if tok, ok := l.consumeLiteral("null", TOKEN_NULL); ok {
			return tok
		}
	}

	// Number: starts with a digit or a minus sign
	if ch >= '0' && ch <= '9' || ch == '-' {
		return l.consumeNumber()
	}

	// Anything else is unexpected input — skip it and signal with empty token
	l.advance()
	return Token{Type: TOKEN_EOF, Value: string(ch)}
}

func main() {
	// input := `{"name": "Rishav", "age": 25, "active": true, "score": -3.14, "tags": ["go", "rust"]}`
	// lexer := NewLexer(input)

	// for {
	// 	tok := lexer.GetNextToken()
	// 	fmt.Printf("%-14s → %q\n", tokenName(tok.Type), tok.Value)
	// 	if tok.Type == TOKEN_EOF {
	// 		break
	// 	}
	// }

	input := `{
		"name": "Rishav",
		"age": 25,
		"active": true,
		"score": -3.14,
		"address": {
			"city": "Nagpur",
			"pin": 440001
		},
		"tags": ["go", "rust", null]
	}`

	parser := NewParser(input)

	value, err := parser.ParseValue()
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	// Pretty-print the resulting Go structure
	printValue(value, 0)
}

// tokenName is just a helper to print readable names instead of raw ints.
func tokenName(t TokenType) string {
	names := map[TokenType]string{
		TOKEN_EOF:      "EOF",
		TOKEN_LBRACE:   "LBRACE",
		TOKEN_RBRACE:   "RBRACE",
		TOKEN_LBRACKET: "LBRACKET",
		TOKEN_RBRACKET: "RBRACKET",
		TOKEN_COMMA:    "COMMA",
		TOKEN_COLON:    "COLON",
		TOKEN_STRING:   "STRING",
		TOKEN_NUMBER:   "NUMBER",
		TOKEN_TRUE:     "TRUE",
		TOKEN_FALSE:    "FALSE",
		TOKEN_NULL:     "NULL",
	}
	return names[t]
}
