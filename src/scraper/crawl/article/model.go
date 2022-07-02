package article

import (
	"fmt"
	"time"
)

type Article struct {
	title   string `json:"title"`
	content string `json:"content"`
	date    string `json:"date"`
}

func NewArticle(title string, content string) *Article {

	timeObject := time.Now()
	year, month, day := timeObject.Date()
	date := fmt.Sprintf("%v-%v-%v", year, int(month), day)

	return &Article{
		title:   title,
		content: content,
		date:    date,
	}
}
