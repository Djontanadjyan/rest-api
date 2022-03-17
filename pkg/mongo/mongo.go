package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var collection *mongo.Collection

func New(url string, db string, col string) *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)
	err = client.Connect(ctx)

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	collection = client.Database(db).Collection(col)

	return collection
}
