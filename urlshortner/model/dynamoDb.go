package model

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDbModel struct {
	dynamoClient *dynamodb.Client
	TableName    string
}

func CreateDynamoDbModel(TableName string) (*DynamoDbModel, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("localhost"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "abcd", SecretAccessKey: "a1b2c3", SessionToken: "",
				Source: "Mock credentials used above for local instance",
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	c := DynamoDbModel{dynamoClient: dynamodb.NewFromConfig(cfg), TableName: TableName}
	return &c, nil
}

// Converts to dynamo keys
func (D *DynamoDbModel) GetDynamoKeys(url UrlModel) map[string]types.AttributeValue {
	EncodedUrl, err := attributevalue.Marshal(url.EncodedUrl)
	if err != nil {
		panic(err)
	}
	Owner, err := attributevalue.Marshal(url.Owner)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"encodedURL": EncodedUrl, "userId": Owner}
}
func (D *DynamoDbModel) Get(ctx context.Context, encodedUrl string, Owner string) (UrlModel, error) {
	var urlR UrlModel
	resp, err := D.dynamoClient.GetItem(context.TODO(), &dynamodb.GetItemInput{Key: D.GetDynamoKeys(UrlModel{EncodedUrl: encodedUrl, Owner: Owner}), TableName: aws.String(D.TableName)})
	if err != nil {
		log.Printf("Cannot Get Item: %v\n", err)

	} else {
		err = attributevalue.UnmarshalMap(resp.Item, &urlR)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return urlR, nil
}
func (D *DynamoDbModel) Set(ctx context.Context, url UrlModel) error {
	item, _ := attributevalue.MarshalMap(url)
	_, err := D.dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(D.TableName), Item: item,
	})
	if err != nil {
		return err
	}
	return nil
}
func (D *DynamoDbModel) Delete(encodedUrl string) error { return nil }
func (D *DynamoDbModel) ListAllKeys()                   {}
