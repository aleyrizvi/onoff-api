package anagram

import "testing"

type testCase struct {
	Word     string
	Anagram  string
	Expected bool
}

func TestIsAnagram(t *testing.T) {
	anagram := New()
	cases := []testCase{
		{
			Word:     "table",
			Anagram:  "batle",
			Expected: true,
		},
		{
			Word:     "tallinn",
			Anagram:  "france",
			Expected: false,
		},
		{
			Word:     "Prefect",
			Anagram:  "perfect",
			Expected: true,
		},
	}

	for _, c := range cases {
		got := anagram.IsAnagram(c.Word, c.Anagram)
		if c.Expected != got {
			t.Errorf("Expected %t but got %t : %v", c.Expected, got, c)
		}
	}
}
