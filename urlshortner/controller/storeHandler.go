package controller

import (
	"context"
	"fmt"
	"shubham/urlShortner/model"
)

type controller struct {
	controller model.UrlStorageDB
}

func CreateController(dbName string) *controller {
	fmt.Println(dbName)
	db, err := CreateStoreHandlerControllerModel(dbName)
	if err != nil {
		fmt.Println("Error in storeController", err)
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
