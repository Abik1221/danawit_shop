package database

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() mongo.Client {
	db, err := mongo.NewClient(options.Client().applyURI("MONGODB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	var ctx, cancel = context.WithTimeout(context.Bckground(), 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

var Client *mongo.Client = DbInstance()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Dana").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Dana").Collection(collectionName)
	return collection
}
