package api

import (
	"github.com/kataras/iris/context"
)

// EchoHandler to be used as Handler for ECHO API
func EchoHandler(ctx context.Context) {
	message := ctx.URLParam("message")
	ctx.Text(message)
}
