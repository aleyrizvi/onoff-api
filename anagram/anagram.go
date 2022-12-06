package anagram

import "strings"

type Anagram struct{}

func New() *Anagram {
	return &Anagram{}
}

func (a *Anagram) IsAnagram(word string, anagram string) bool {

	if len(word) != len(anagram) {
		return false
	}

	anagramMap := make(map[string]int)

	// iterating first word
	for i := 0; i < len(word); i++ {
		anagramMap[strings.ToLower(string(word[i]))]++
	}

	// iterating second word to check if character exists in map, reduce the count
	for i := 0; i < len(anagram); i++ {
		anagramMap[strings.ToLower(string(anagram[i]))]--
	}

	// verify the map
	for i := 0; i < len(word); i++ {
		if anagramMap[strings.ToLower(string(word[i]))] != 0 {
			return false
		}
	}

	return true
}
