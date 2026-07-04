package handler

import (
	"calculator_cli/operations"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Result float64 `json:"result"`
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		http.Error(w, "%d invalid a", http.StatusBadRequest)
		return
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		http.Error(w, "%d invalid b", http.StatusBadRequest)
		return
	}
	result := operations.Add(a, b)
	response := Response{
		Result: result,
	}
	json.NewEncoder(w).Encode(response)
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Fprintf(w, " %d invalid a", http.StatusBadRequest)
		return
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Fprintf(w, " %d invalid a", http.StatusBadRequest)
		return
	}
	result := operations.Substract(a, b)
	response := Response{
		Result: result,
	}
	json.NewEncoder(w).Encode(response)
}

func MulHandler(w http.ResponseWriter, r *http.Request) {

}

func DivHandler(w http.ResponseWriter, r *http.Request) {

}
