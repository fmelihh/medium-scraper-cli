package crawlArticle

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"medium-scraper-cli/src/scraper/crawl"
	"medium-scraper-cli/src/scraper/db"
	"time"
)

type Article struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title,omitempty"`
	Content string             `bson:"content,omitempty"`
	Date    string             `bson:"date,omitempty"`
}

func NewArticle(title string, content string) *Article {

	timeObject := time.Now()
	year, month, day := timeObject.Date()
	date := fmt.Sprintf("%v-%v-%v", year, int(month), day)

	return &Article{
		Title:   title,
		Content: content,
		Date:    date,
	}
}

func (article *Article) SaveOnMongoDB() (string, error) {

	client, ctx, cancel, err := db.ConnectMongo()
	if err != nil {
		return "", nil
	}

	defer db.CloseMongo(client, ctx, cancel)

	collection := client.Database(crawl.DatabaseName).Collection(crawl.ArticleCollection)
	result, err := collection.InsertOne(ctx, article)
	if err != nil {
		return "FAIL", err
	}

	message := fmt.Sprintf("Article Inserted: %v", result.InsertedID)

	return message, nil
}
