package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	} {
		{
			input: 	"	hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "hello   there  ",
			expected: []string{"hello", "there"},
		},
		{
			input: "HeLLo   STEVE  ",
			expected: []string{"hello", "steve"},
		},
		{
			input: "   ",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("CleanInput function generated an output of length not matching the expected")
			t.Fail()
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("CleanInput function generated an output with words not matching the expected")
				t.Fail()
			}
		}	
	}
}

