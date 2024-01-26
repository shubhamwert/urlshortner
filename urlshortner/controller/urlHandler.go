package controller

import (
	"context"
	"fmt"
	"shubham/urlShortner/model"
)

type UrlController struct {
	db *controller
}

func CreateUrlController() UrlController {
	return UrlController{
		db: CreateController(),
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

	urlObject, err := u.db.Get(ctx, url)
	if err != nil {
		fmt.Println("ERROR ", err)

		return "", nil
	}
	return urlObject.OriginalUrl, nil
}
