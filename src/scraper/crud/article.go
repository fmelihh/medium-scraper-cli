package crud

import (
	"go.mongodb.org/mongo-driver/bson"
	"medium-scraper-cli/src/scraper/crawl"
	"medium-scraper-cli/src/scraper/crawl/crawlArticle"
	"medium-scraper-cli/src/scraper/db"
)

func ListArticles(args bson.M) (*[]crawlArticle.Article, error) {

	articles := make([]crawlArticle.Article, 5)

	client, ctx, cancel, err := db.ConnectMongo()
	if err != nil {
		return nil, nil
	}

	defer db.CloseMongo(client, ctx, cancel)
	collection := client.Database(crawl.DatabaseName).Collection(crawl.ArticleCollection)

	results, err := collection.Find(ctx, args)
	if err != nil {
		return nil, err
	}

	if err := results.All(ctx, &articles); err != nil {
		return nil, err
	}

	return &articles, nil
}
