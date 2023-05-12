package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alwindoss/morse"
)

// Solution
func DecodeMorse(morseCode string) string {
	h := morse.NewHacker()

	words := strings.Split(morseCode, "   ")

	for i := 0; i < len(words); i++ {
		res, err := h.Decode(strings.NewReader(words[i]))
		if err != nil {
			return ""
		}
		words[i] = string(res)
	}

	return strings.Join(words, " ")
}

func TestDecodeMorse(t *testing.T) {
	tests := []struct {
		a    string
		want string
	}{
		{"  .... . -.--   .--- ..- -.. .", "HEY JUDE"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := DecodeMorse(tt.a)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}

		})
	}
}
