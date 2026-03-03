package main

import "fmt"

func IsValidSudoku(board [][]byte) bool {
	// byte or rune are used for ascii representation
	// '5' is not int 5 but 53 in ascii val

	//usually these kind of problems uses strings as the input
	//converting them to ascii is easier and memory efficient

	//byte: uses 1 byte (8 bits)
	//int : uses 8 bytes (64 bit) on 64 bit system

	//using byte instead of int here is because the input size
	//is fixed i.e. 9 which is too small and allocating 8 bytes of
	//memory just for len 9 matrix is like using a truct to
	//transfer an envelope

	// char byte val   calculation result
	// '1'     49 	           49 - 49    0  (index)
	// '2'	   50              50 - 49    1
	// and so on

	// In Go lang a string is essentially a read-only slice
	// of bytes. Converting a string input into a [][]byte is nearly instantaneous
	// whereas converting it to [][]int requires nested loops and manual parsing.

	//Summary
	// We use byte because
	//1. Directness: It matches how text data is stored
	//2. Speed: It allows for 'offset math' instead of function calls strconv.Atoi
	//3. Parsimony; It uses 8x less memory than integer

	rows := make([][10]bool, 9)
	cols := make([][10]bool, 9)
	boxes := make([][10]bool, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				//calculate box ind: 0 to 8
				boxInd := (i/3)*3 + (j / 3)
				//check for existing presence
				if rows[i][num] || cols[j][num] || boxes[boxInd][num] {
					return false
				}
				//mark as seen
				rows[i][num] = true
				cols[j][num] = true
				boxes[boxInd][num] = true
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	isValid := IsValidSudoku(board)
	if isValid {
		fmt.Println("The sudoku is valid")
	} else {
		fmt.Println("The sudoku is not valid")
	}
}
