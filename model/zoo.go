package model

type Animal struct {
	Name Zoo `json:"name"`
}

type Zoo string

const (
	Cat   Zoo = "cat"
	Dog   Zoo = "dog"
	Woman Zoo = "woman"
)
