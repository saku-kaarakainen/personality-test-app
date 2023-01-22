package entity

type Result struct {
	Id                    string   `json:"id"`
	Label                 string   `json:"label"`
	DescriptionParagraphs []string `json:"description_paragraphs"`
}
