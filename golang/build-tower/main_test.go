package main

import (
	"fmt"
	"strings"
	"testing"
)

// Solution
func TowerBuilder(nFloors int) []string {
	a := make([]string, nFloors)
	width := (nFloors-1)*2 + 1

	for i := 0; i < nFloors; i++ {
		s := []rune(strings.Repeat("*", width))

		for j := 0; j < width/2-i; j++ {
			s[j] = ' '
		}

		for j := width - 1; j > width/2+i; j-- {
			s[j] = ' '
		}

		a[i] = string(s)
	}

	return a
}

func TestTowerBuilder(t *testing.T) {
	tests := []struct {
		a    int
		want []string
	}{
		{1, []string{"*"}},
		{2, []string{" * ", "***"}},
		{3, []string{"  *  ", " *** ", "*****"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %v/Expect: %v", tt.a, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := TowerBuilder(tt.a)
			if !compare(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}

		})
	}
}

func compare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
