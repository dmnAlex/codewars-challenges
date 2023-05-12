package main

import (
	"fmt"
	"testing"
)

// Solution
func InAscOrder(numbers []int) bool {
	if len(numbers) == 0 {
		return true
	}

	last := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if last > numbers[i] {
			return false
		}
		last = numbers[i]
	}

	return true
}

func TestDecodeMorse(t *testing.T) {
	tests := []struct {
		a    []int
		want bool
	}{
		{[]int{1, 2, 4, 7, 19}, true},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{1, 6, 10, 18, 2, 4, 20}, false},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %v/Expect: %v", tt.a, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := InAscOrder(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}

		})
	}
}
