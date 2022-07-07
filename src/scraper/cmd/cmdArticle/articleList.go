package cmdArticle

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"medium-scraper-cli/src/crud"
)

func ArticleList(params string) {
	a := make(bson.M, 0)
	a["title"] = "IMPLEMENT HEXAGONAL ARCHITECTURE"

	articles, err := crud.ListArticles(a)
	if err != nil {
		fmt.Println(err)
	}

	for i, article := range *articles {
		fmt.Println(i, article.Title, article.Date, article.ID)
	}

}
