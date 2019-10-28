package main

import (
	"os"

	"github.com/iris-contrib/cloud-native-go/api"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.Get("/", index)

	apiGroup := app.Party("/api")
	{
		apiGroup.Get("/echo", api.EchoHandler)
		apiGroup.Get("/hello", api.HelloHandler)
		apiGroup.Get("/books", api.AllBooksHandler)
		apiGroup.Get("/books/{isbn:string}", api.GetBookHandler)
		apiGroup.Post("/books", api.CreateBookHandler)
		apiGroup.Put("/books/{isbn:string}", api.UpdateBookHandler)
		apiGroup.Delete("/books/{isbn:string}", api.DeleteBookHandler)

	}

	app.Run(iris.Addr(":" + port()))
}

func port() string {
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
