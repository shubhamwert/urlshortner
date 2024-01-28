package controller

import (
	"context"
	"fmt"
	"shubham/urlShortner/model"
)

type controller struct {
	controller model.UrlStorageDB
}

func CreateController() *controller {
	// db := model.CreateinMemoryStorage()
	// db := model.CreateredisStorage()
	// db := model.CreatemongoStorage("mongodb://admin:password@localhost:27017/", "UrlDb", "ShortenedUrl")
	db, err := model.CreateDynamoDbModel("UrlDb")
	if err != nil {
		fmt.Println("Error in storeController")
	}
	controller := &controller{
		controller: db,
	}
	return controller
}

func (c *controller) Get(ctx context.Context, url string, owner string) (model.UrlModel, error) {
	return c.controller.Get(ctx, url, owner)

}
func (c *controller) Set(ctx context.Context, url model.UrlModel) error {
	return c.controller.Set(ctx, url)

}

func (c *controller) Delete() {}

func (c *controller) Test() {
	ctx := context.Background()
	m, err := model.CreateUrlModel("12", "a", "b", "test")
	if err != nil {
		fmt.Println("Mem not working")
	}
	c.controller.Set(
		ctx, *m,
	)
	v, e := c.controller.Get(ctx, "b", "test")
	fmt.Println(v, e)
	v2, e2 := c.controller.Get(ctx, "a", "test")
	fmt.Println(v2, e2)

}
