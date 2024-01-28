package controller

import (
	"context"
	"fmt"
	"shubham/urlShortner/model"
)

type cacheController struct {
	cacheClient model.UrlStorageDB
}

func createCacheController() *cacheController {
	c := model.CreateredisStorage("localhost:6379", "", 0)
	return &cacheController{
		cacheClient: c,
	}

}

func (c *cacheController) Get(ctx context.Context, url string, owner string) (model.UrlModel, error) {

	u, err := c.cacheClient.Get(ctx, url, owner)

	if err != nil {
		return model.UrlModel{}, err
	}
	if ctx.Err() == context.DeadlineExceeded {

		return model.UrlModel{}, fmt.Errorf("ctx deadline exceded")
	}
	return u, nil
}
func (c *cacheController) Set(ctx context.Context, url model.UrlModel) error {
	return c.cacheClient.Set(ctx, url)

}
