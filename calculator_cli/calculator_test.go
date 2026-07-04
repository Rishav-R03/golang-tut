package main

import (
	"calculator_cli/operations"
	"testing"
)

func TestAdd(t *testing.T) {
	got := operations.Add(10, 30)
	want := 40.0

	if got != want {
		t.Errorf("got %v want %v \n", got, want)
	}
}

func TestSubtraction(t *testing.T) {
	got := operations.Substract(30, 10)
	want := 20.0
	if got != want {
		t.Errorf("got %v want %v \n", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := operations.Multiply(30, 10)
	want := 300.0
	if got != want {
		t.Errorf("got %v want %v \n", got, want)
	}
}

func TestDivision(t *testing.T) {
	got, err := operations.Divide(30, 10)
	if err != nil {
		t.Errorf("divide by zero error: %s", err)
		return
	}
	want := 3.0
	if got != want {
		t.Errorf("got %v want %v \n", got, want)
	}
}
