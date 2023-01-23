package entity

type Result struct {
	Id                    string   `json:"id"`
	Label                 string   `json:"label"`
	Score                 int      `json:"score"`
	DescriptionParagraphs []string `json:"description_paragraphs"`
}
