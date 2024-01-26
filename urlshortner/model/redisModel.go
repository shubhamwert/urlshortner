package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type redisStorage struct {
	redisClient *redis.Client
}

func CreateredisStorage() *redisStorage {
	redisClient := &redisStorage{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}

	return redisClient

}

func (M *redisStorage) Set(ctx context.Context, url UrlModel) error {
	urlJson, err := json.Marshal(url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return M.redisClient.Set(ctx, url.EncodedUrl, urlJson, 0).Err()
}

func (M *redisStorage) Get(ctx context.Context, url string) (UrlModel, error) {

	val, err := M.redisClient.Get(ctx, url).Result()

	if err != nil {
		fmt.Println("error at get url", url, " err ", err)
		return UrlModel{}, err
	}
	result := &UrlModel{}
	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		fmt.Println("error at unmarshal ", err)
	}
	return *result, err
}

func (M *redisStorage) ListAllKeys() {
	fmt.Println("Getting Keys")
	var cursor uint64
	ctx := context.Background()
	for {
		var keys []string
		var err error
		keys, cursor, err = M.redisClient.Scan(ctx, cursor, "*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
			s, err := M.redisClient.Get(ctx, key).Result()
			if err != nil {
				fmt.Println("error ", err)
			}
			fmt.Println("val", s)
		}

		if cursor == 0 { // no more keys
			break
		}
	}

}

func (M *redisStorage) Delete(url string) error {
	return nil
	// _, ok := M.Url[url]
	// if !ok {
	// 	return errors.New("Encoding do not exists")
	// }

	// delete(M.Url, url)
	// return nil
}
