package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello","world"},
		},
		{
			input: "fahad ahmed ",
			expected: []string{"fahad", "ahmed"},
		},
		{
			input: " Golang",
			expected: []string{"golang"},
		},
	}

	for _,c := range cases {
		actual := cleanInput((c.input))
		if len(actual) != len(c.expected) {
			t.Errorf("Length of slice is not equal")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected word is not equal to actual word")
			}
		}
	}
}

