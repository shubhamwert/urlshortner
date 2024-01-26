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
	c := model.CreateredisStorage()
	return &cacheController{
		cacheClient: c,
	}

}

func (c *cacheController) Get(ctx context.Context, url string) (model.UrlModel, error) {
	u, err := c.cacheClient.Get(ctx, url)

	if err != nil {
		return model.UrlModel{}, err
	}
	if ctx.Err() == context.DeadlineExceeded {

		return model.UrlModel{}, fmt.Errorf("ctx deadline excedded")
	}
	return u, nil
}
func (c *cacheController) Set(ctx context.Context, url model.UrlModel) error {
	return c.cacheClient.Set(ctx, url)

}
