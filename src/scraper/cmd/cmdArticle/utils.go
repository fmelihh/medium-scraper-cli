package cmdArticle

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func ParamsToBsonM(params string) (bson.M, error) {
	bsonObject := make(bson.M, 0)

	splittedParams := strings.Split(params, ",")

	for _, param := range splittedParams {
		values := strings.Split(param, "=")
		if len(values) > 2 {
			panic("Incorrect parameter")
		}

		title := strings.ToLower(values[0])
		value := values[1]

		bsonObject[title] = value
	}

	return bsonObject, nil

}
