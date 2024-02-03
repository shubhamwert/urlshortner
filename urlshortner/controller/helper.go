package controller

import (
	"errors"
	"shubham/urlShortner/model"

	"github.com/spf13/viper"
)

func CreateStoreHandlerControllerModel(modelName string) (model.UrlStorageDB, error) {

	switch modelName {
	case "mongodb":
		return model.CreatemongoStorage(viper.GetString("mongo.url"), viper.GetString("mongo.db"), viper.GetString("mongo.Collection")), nil
	case "dynamodb":
		return model.CreateDynamoDbModel(viper.GetString("dynamodb.tablename"))
	// case "inmemory":
	// 	return model.CreateinMemoryStorage(), nil
	case "redis":
		return model.CreateredisStorage(viper.GetString("redis.addr"), viper.GetString("redis.pass"), viper.GetInt("DefaultDb")), nil
	default:
		return nil, errors.New("no db found")
	}

}
