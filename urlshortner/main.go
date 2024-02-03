package main

import (
	"flag"
	"fmt"
	"log"
	"shubham/urlShortner/controller"

	"github.com/spf13/viper"
)

func main() {
	configPath := flag.String("configPath", "./configs", "Provides the config")
	configName := flag.String("configName", "config.json", "Provides the config")
	flag.Parse()
	fmt.Println(*configPath, "  ", *configName)
	CreateConfig(*configPath, *configName)
	log.Printf("Running Url Shortner:\n Storage Db: %s \n Cache: %s", viper.GetString("storagedb"), viper.GetString("cachedb"))
	app := CreateApp()
	handler := controller.CreateUrlController(viper.GetString("storagedb"), viper.GetString("cachedb"))
	InitializeApp(app, &handler)
	RunApp(app)
}
