package main

import "shubham/urlShortner/controller"

func main() {
	app := CreateApp()
	handler := controller.CreateUrlController()
	InitializeApp(app, &handler)
	handler.Test()
	RunApp(app)
}
