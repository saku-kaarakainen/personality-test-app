package entity

type Answer struct {
	Id    string `json:"id"`
	Score [2]int `json:"score"`
	Label string `json:"answer_label"`
}
