package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hellO woRld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		// add more cases here
	}

	for _, tc := range cases {
		result := cleanInput(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf("length of slices do not match.  Expected: %d  Actual:  %d", len(tc.expected), len(result))
			t.Fail()
		} else {
			for i, rw := range result {
				ew := tc.expected[i]
				if rw != ew {
					t.Errorf("words do not match. Expected: %s  Actual: %s", ew, rw)
					t.Fail()
				}
			}
		}

	}
}
