package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	err := M.Coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "EncodedUrl", Value: url}}).Decode(&urlModel)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the url %s %s\n", url, err)
		return urlModel, nil
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
	//find records
	//pass these options to the Find method
	findOptions := options.Find()
	//Set the limit of the number of record to find
	findOptions.SetLimit(5)
	//Define an array in which you can store the decoded documents
	var results []UrlModel

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	cur, err := M.Coll.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Finding multiple documents returns a cursor
	//Iterate through the cursor allows us to decode documents one at a time

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem UrlModel
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", results)

}

func (M *mongoStorage) Delete(url string) error {
	return nil

}
