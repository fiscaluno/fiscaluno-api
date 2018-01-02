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

func Handle(app *iris.Application) {
	usersRoutes := app.Party("/users", userMiddleware)
	{
		usersRoutes.Get("/id", teste)
	}
}
