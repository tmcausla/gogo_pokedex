package main
import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   Go   is   FUN   ",
			expected: []string{"go", "is", "fun"},
		},
		{
			input:    "MiXeD CaSe WoRdS",
			expected: []string{"mixed", "case", "words"},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "     ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("hashtag epic split fail")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] = %q, want %q", 
					c.input, i, word, expectedWord)
			}
		}
	}
}