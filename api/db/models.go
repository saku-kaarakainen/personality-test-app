package db

// questions
type Answer struct {
	Id    string `json:"id"`
	Label string `json:"question_label"`
}

type Question struct {
	Id      string   `json:"id"`
	Label   string   `json:"question_label"`
	Answers []Answer `json:"answers"`
}

// results
type Result struct {
	Id                    string   `json:"id"`
	Label                 string   `json:"question_label"`
	DescriptionParagraphs []string `json:"description_paragraphs"`
}
