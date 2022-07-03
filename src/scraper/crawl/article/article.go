package article

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

var title string
var content string

func ScrapeArticle(url string) (string, error) {
	c := colly.NewCollector()

	c.SetRequestTimeout(120 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\nvisiting.. %v \n", r.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Printf("\ngot response from %v \n", response.Request.URL)
	})

	c.OnHTML("article", func(e *colly.HTMLElement) {
		title = TitleFromURL(url)
		content = e.Text
	})

	err := c.Visit(url)
	if err != nil {
		return "FAIL", err
	}

	article := NewArticle(title, content)
	message, err := article.SaveOnMongoDB()
	if err != nil {
		return message, err
	}

	return "SUCCESS", nil
}
