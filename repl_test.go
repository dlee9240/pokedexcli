package main

import "testing"

func TestCleanInput(t *testing.T) {
	//...
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{ //second case
			input:    " Daniel Lee is the best  person",
			expected: []string{"daniel", "lee", "is", "the", "best", "person"},
		},
		{ //third case
			input:    " Test CASE 3",
			expected: []string{"test", "case", "3"},
		},
	}

	//add more cases here

	for _, c := range cases {
		actual := cleanInput(c.input)
		//check the length of the actual slice against the expected slice:
		//if they don't match, use t.Errorf to print an error message and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %q, want %q", c.input, actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			//need to finish the function
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("got %q, expected %q", word, expectedWord)
			}

		}
	}
}

/*
This is boot.dev solution...
package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
*/
