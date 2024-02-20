//go:build lambda_dynamo

package controller

import (
	"fmt"
	"shubham/urlShortner/model"

	"github.com/spf13/viper"
)

func CreateStoreHandlerControllerModel(modelName string) (model.UrlStorageDB, error) {
	fmt.Println("Getting DynamoDb connection")
	return model.CreateDynamoDbModel(viper.GetString("dynamodb.tablename"))

}
