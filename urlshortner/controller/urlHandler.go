package controller

import (
	"context"
	"fmt"
	"shubham/urlShortner/model"
	"shubham/urlShortner/repo"
	"time"

	"github.com/spf13/viper"
)

type UrlController struct {
	db    *controller
	cache *cacheController
}

var TIMEOUT time.Duration

func CreateUrlController(urlStoreName string, urlCacheName string) UrlController {
	var cache *cacheController
	TIMEOUT = viper.GetDuration("cacheTimeout")
	if viper.GetBool("cacheEnabled") {
		cache = createCacheController(urlCacheName)
	} else {
		cache = nil
	}
	return UrlController{
		db:    CreateController(urlStoreName),
		cache: cache,
	}
}

func (u *UrlController) Shorten(ctx context.Context, url string, owner string) (string, error) {

	urlObject := model.UrlModel{}
	urlObject.OriginalUrl = url
	urlObject.EncodedUrl = repo.CreateShortUrl(url, 5)
	urlObject.Owner = owner
	fmt.Println("Shorten ", urlObject)
	err := u.db.Set(ctx, urlObject)
	if err != nil {
		fmt.Println(err)
	}
	return urlObject.EncodedUrl, nil
}

func (u *UrlController) GetUrl(ctx context.Context, url string, owner string) (string, error) {
	// Try adding cachce
	CacheCtx, cacheCancel := context.WithTimeout(ctx, 3*time.Second)
	defer cacheCancel()
	cacheChannel := make(chan model.UrlModel, 1)
	if viper.GetBool("cacheEnabled") {
		go func() {
			urlObject, err := u.cache.Get(CacheCtx, url, owner)
			if err != nil {
				fmt.Println("Error getting from cache")
				return
			}
			cacheChannel <- urlObject
		}()
	}
	var urlObject model.UrlModel
	select {
	case urlObject = <-cacheChannel:
		fmt.Println("Cache result:", urlObject)
	case <-time.After(TIMEOUT * time.Second):
		var err error
		urlObject, err = u.db.Get(ctx, url, owner)
		if err != nil {
			fmt.Println("ERROR ", err, ctx)

			return "", nil
		}
	}
	if viper.GetBool("cacheEnabled") {
		go u.cache.Set(context.TODO(), urlObject)
	}
	return urlObject.OriginalUrl, nil
}
