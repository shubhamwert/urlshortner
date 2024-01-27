package controller

import (
	"context"
	"fmt"
	"shubham/urlShortner/model"
	"time"
)

type UrlController struct {
	db    *controller
	cache *cacheController
}

func CreateUrlController() UrlController {
	return UrlController{
		db:    CreateController(),
		cache: createCacheController(),
	}
}

func (u *UrlController) Shorten(ctx context.Context, url string) (string, error) {

	urlObject := model.UrlModel{}
	urlObject.OriginalUrl = url
	urlObject.EncodedUrl = url[:8]
	err := u.db.Set(ctx, urlObject)
	if err != nil {
		fmt.Println(err)
	}
	return urlObject.EncodedUrl, nil
}

func (u *UrlController) GetUrl(ctx context.Context, url string) (string, error) {
	// Try adding cachce
	CacheCtx, cacheErr := context.WithTimeout(ctx, 3*time.Second)
	defer cacheErr()
	cacheChannel := make(chan model.UrlModel, 1)
	go func() {
		urlObject, err := u.cache.Get(CacheCtx, url)
		if err != nil {
			fmt.Println("Error getting from cache")
			return
		}
		cacheChannel <- urlObject
	}()
	var urlObject model.UrlModel
	select {
	case urlObject = <-cacheChannel:
		fmt.Println("Cache result:", urlObject)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
		var err error
		urlObject, err = u.db.Get(ctx, url)
		if err != nil {
			fmt.Println("ERROR ", err)

			return "", nil
		}
	}
	fmt.Println("URL: ", urlObject.OriginalUrl)

	go u.cache.Set(context.TODO(), urlObject)
	return urlObject.OriginalUrl, nil
}
