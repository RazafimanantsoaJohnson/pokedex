package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    johnson LEGRAND",
			expected: []string{"johnson", "legrand"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The result of clean input: \n %v \nis not the same as the one we expected: \n %v", actual, c.expected)
			return
		}

		for i, k := range c.expected {
			if actual[i] != k {
				t.Errorf("The result of clean input: \n %v \nis not the same as the one we expected: \n %v", actual, c.expected)
				return
			}
		}
	}
}
