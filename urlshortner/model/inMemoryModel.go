package model

import (
	"context"
	"errors"
	"fmt"
)

type inMemoryStorage struct {
	Url map[string]UrlModel `json:"url"`
}

func CreateinMemoryStorage() *inMemoryStorage {
	return &inMemoryStorage{Url: make(map[string]UrlModel)}

}

func (M *inMemoryStorage) Set(ctx context.Context, url UrlModel) error {

	_, ok := M.Url[url.EncodedUrl]
	fmt.Println("Used String for encoding ", url.EncodedUrl)
	if !ok {
		M.Url[url.EncodedUrl] = url
		fmt.Println("Listing all keys")

		M.ListAllKeys()
		return nil
	}
	fmt.Println("Key already exists Listing all keys")
	M.ListAllKeys()

	return errors.New("Encoding Already exists")
}

func (M *inMemoryStorage) Get(ctx context.Context, url string) (UrlModel, error) {
	M.ListAllKeys()

	val, ok := M.Url[url]
	fmt.Println("value is ", val)
	if !ok {
		fmt.Println("Listing all keys")
		M.ListAllKeys()
		return UrlModel{}, errors.New("Encoding do not exists at InMemory Db")
	}

	return val, nil
}

func (M *inMemoryStorage) ListAllKeys() {
	for k, v := range M.Url {
		fmt.Println(k, "                  ", v)
	}
}

func (M *inMemoryStorage) Delete(url string) error {

	_, ok := M.Url[url]
	if !ok {
		return errors.New("Encoding do not exists")
	}

	delete(M.Url, url)
	return nil
}
