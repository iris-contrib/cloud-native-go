package api

import "github.com/kataras/iris/v12"

// EchoHandler to be used as Handler for ECHO API
func EchoHandler(ctx iris.Context) {
	message := ctx.URLParam("message")
	ctx.Text(message)
}
