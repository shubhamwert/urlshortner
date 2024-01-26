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
	db := model.CreatemongoStorage()
	controller := &controller{
		controller: db,
	}
	return controller
}

func (c *controller) Get(ctx context.Context, url string) (model.UrlModel, error) {
	return c.controller.Get(ctx, url)

}
func (c *controller) Set(ctx context.Context, url model.UrlModel) error {
	return c.controller.Set(ctx, url)

}

func (c *controller) Delete() {}

func (c *controller) Test() {
	ctx := context.Background()
	m, err := model.CreateUrlModel("1", "a", "b")
	if err != nil {
		fmt.Println("Mem not working")
	}
	c.controller.Set(
		ctx, *m,
	)
	v, e := c.controller.Get(ctx, "b")
	fmt.Println(v, e)
	v2, e2 := c.controller.Get(ctx, "a")
	fmt.Println(v2, e2)

}
