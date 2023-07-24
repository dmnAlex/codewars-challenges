package main

import (
	"fmt"
	"strings"
	"testing"
)

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Peek() (interface{}, bool) {
	size := len(*s)
	if size > 0 {
		return (*s)[size-1], true
	} else {
		return nil, false
	}
}

func (s *Stack) Pop() (interface{}, bool) {
	size := len(*s)
	if size > 0 {
		item := (*s)[size-1]
		*s = (*s)[:(size - 1)]
		return item, true
	} else {
		return nil, false
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func ValidBraces(str string) bool {
	stack := Stack{}
	const leftBrackets = "({["
	const rightBrackets = ")}]"

	for i := range str {
		if index := strings.IndexByte(leftBrackets, str[i]); index != -1 {
			stack.Push(leftBrackets[index])
		} else if index := strings.IndexByte(rightBrackets, str[i]); index != -1 {
			item, ok := stack.Pop()
			if ok != true || item != leftBrackets[index] {
				return false
			}
		} else {
			return false
		}
	}

	return stack.IsEmpty()
}
func TestValidBraces(t *testing.T) {
	tests := []struct {
		a    string
		want bool
	}{
		{"(){}[]", true},
		{"([{}])", true},
		{"(}", false},
		{"[(])", false},
		{"[({)](]", false},
		{"]", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %v/Expect: %v", tt.a, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := ValidBraces(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
