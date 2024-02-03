package main

import (
	"log"
	"shubham/urlShortner/controller"

	"github.com/spf13/viper"
)

func main() {
	CreateConfig("./configs", "config.json")
	log.Printf("Running Url Shortner:\n Storage Db: %s \n Cache: %s", viper.GetString("storagedb"), viper.GetString("cachedb"))
	app := CreateApp()
	handler := controller.CreateUrlController(viper.GetString("storagedb"), viper.GetString("cachedb"))
	InitializeApp(app, &handler)
	RunApp(app)
}
