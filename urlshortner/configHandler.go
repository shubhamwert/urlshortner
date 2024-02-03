package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func CreateConfig(path string, config string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(config)
	viper.SetConfigType("json")
	viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
