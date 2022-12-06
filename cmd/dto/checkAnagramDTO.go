package dto

type CheckAnagramRequestBody struct {
	Word    string `json:"word" validate:"required"`
	Anagram string `json:"anagram" validate:"required"`
}

type CheckAnagramResponseBody struct {
	IsAnagram bool `json:"isAnagram"`
}
