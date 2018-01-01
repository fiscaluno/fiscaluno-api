package routes

import (
	"github.com/kataras/iris"
)

func teste(ctx iris.Context) {
	ctx.Writef("teste")
}

func userMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

var app = iris.New()

// Iris application containing routes
func Router() *iris.Application {
	return app
}
