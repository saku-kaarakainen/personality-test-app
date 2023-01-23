package entity

type Question struct {
	Id          string   `json:"id"`
	Text        string   `json:"question_text"`
	Description string   `json:"question_description"`
	Answers     []Answer `json:"answers"`
}
