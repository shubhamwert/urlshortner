package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStorage struct {
	mongoClient *mongo.Client
	Db          string
	Collection  string
	Coll        *mongo.Collection
}

func CreatemongoStorage() *mongoStorage {
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Error at mongo conn")
	}
	mongoClient := mongoStorage{
		mongoClient: client,
	}
	mongoClient.Db = "UrlDb"
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	mongoClient.Coll = mongoClient.mongoClient.Database(mongoClient.Db).Collection("ShortenedUrl")

	if err != nil {
		log.Fatal(err)
	}

	return &mongoClient

}

func (M *mongoStorage) Set(ctx context.Context, url UrlModel) error {

	result, err := M.Coll.InsertOne(
		context.TODO(),
		url,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return err
}

func (M *mongoStorage) Get(ctx context.Context, url string) (UrlModel, error) {
	var urlModel UrlModel
	err := M.Coll.FindOne(context.TODO(), bson.D{{"url", url}}).Decode(&urlModel)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the url %s\n", url)
		return UrlModel{}, nil
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(urlModel, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return urlModel, nil
}

func (M *mongoStorage) ListAllKeys() {

}

func (M *mongoStorage) Delete(url string) error {
	return nil

}
