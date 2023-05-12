package main

import (
	"fmt"
	"testing"
)

func Multiple3And5(number int) int {
	sum := 0

	for i := 3; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func TestMultiple3And5(t *testing.T) {
	tests := []struct {
		a    int
		want int
	}{
		{10, 23},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Multiple3And5(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}

		})
	}
}
