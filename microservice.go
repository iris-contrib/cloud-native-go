package main

import (
	"os"

	"github.com/iris-contrib/cloud-native-go/api"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/", index)

	apiGroup := app.Party("/api")
	{
		apiGroup.Get("/echo", api.EchoHandler)
		apiGroup.Get("/hello", api.HelloHandler)
		apiBooksGroup := apiGroup.Party("/books")
		{
			apiBooksGroup.Get("/", api.AllBooksHandler)
			apiBooksGroup.Get("/{isbn:string}", api.GetBookHandler)
			apiBooksGroup.Post("/", api.CreateBookHandler)
			apiBooksGroup.Put("/{isbn:string}", api.UpdateBookHandler)
			apiBooksGroup.Delete("/{isbn:string}", api.DeleteBookHandler)
		}
	}

	app.Listen(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if port[0] == ':' {
		port = port[1:]
	}

	return port
}

func index(ctx iris.Context) {
	ctx.Writef("Welcome to Cloud Native")
}
