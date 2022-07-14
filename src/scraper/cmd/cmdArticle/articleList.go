package cmdArticle

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"medium-scraper-cli/src/scraper/crawl/crawlArticle"
	"medium-scraper-cli/src/scraper/crud"
)

func ArticleList(params string) {

	bsonParams, err := ParamsToBsonM(params)
	if err != nil {
		panic(err)
	}

	articles, err := crud.ListArticles(bsonParams)
	if err != nil {
		fmt.Println(err)
	}

	articleList := make([]crawlArticle.Article, len(*articles))
	articleTitles := make([]string, len(*articles))

	for _, article := range *articles {
		articleList = append(articleList, article)
		articleTitles = append(articleTitles, article.Title)
	}

	promt := promptui.Select{
		Label: "Articles",
		Items: articleTitles,
	}

	index, _, err := promt.Run()
	if err != nil {
		fmt.Printf("Promt failed %v\n", err)
		return
	}

	selectedArticle := articleList[index]
	fmt.Printf("You choose %v \n", selectedArticle.Title)

}
