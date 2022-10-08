package model

type GetTitleFilter struct {
	Code string
}

type Title struct {
	Id    int    `json:"id"`
	Code  string `json:"code"`
	Names Names  `json:"names"`
}

type Names struct {
	Ru          string  `json:"ru"`
	En          string  `json:"en"`
	Alternative *string `json:"alternative"`
}
