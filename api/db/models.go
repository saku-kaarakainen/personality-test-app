package db

// questions
type Answer struct {
	Id    string `json:"id"`
	Label string `json:"answer_label"`
}

type Question struct {
	Id          string   `json:"id"`
	Text        string   `json:"question_text"`
	Description string   `json:"question_description"`
	Answers     []Answer `json:"answers"`
}

// results
type Result struct {
	Id                    string   `json:"id"`
	Label                 string   `json:"question_label"`
	DescriptionParagraphs []string `json:"description_paragraphs"`
}
