package main

import (
	"strings"
	"testing"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.Trim(strings.ToLower(text), " "))
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELlo WOrWd",
			expected: []string{"hello", "worwd"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expected_word := c.expected[i]
			t.Log(word, " : ", expected_word)
			if word != expected_word {
				t.Fail()
			}
		}
	}

}
