package cmdArticle

import (
	"fmt"
	"medium-scraper-cli/src/scraper/crawl/crawlArticle"
	"net/url"
	"strings"
)

func ArticleUrl(params string) {

	splitParams := strings.Split(params, ",")

	successCount := len(splitParams)

	for _, param := range splitParams {

		confirmedUrl, err := func(inputUrl string) (string, error) {
			_, err := url.ParseRequestURI(inputUrl)
			return inputUrl, err
		}(param)

		if err != nil {
			successCount = successCount - 1
			fmt.Printf("%v: %v", err, param)
			continue
		}

		message, err := crawlArticle.ScrapeArticle(confirmedUrl)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(message)
	}
}
