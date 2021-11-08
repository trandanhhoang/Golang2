package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	fmt.Println("hoang test")
	if Calculate(2) != 4 {
		t.Error("2+2=4")
	}
}

func TestTable(t *testing.T) {
	var tests = []struct {
		input  int
		expect int
	}{
		// {2, 5},
		{3, 5},
		// {4, 9},
		{5, 7},
		{-1, 1},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expect {
			t.Error(`Test Failed {} input,{} expect, received:{}`, test.input, test.expect, output)
		}
	}
}
