package routes

import (
    "github.com/kataras/iris"
)

// Test Function
func teste(ctx iris.Context) {
    ctx.Writef("teste")
}

// Defualt User Middleware
func userMiddleware(ctx iris.Context) {
    ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
    ctx.Next()
}

var app = iris.New()

// Iris application containing routes
func Router() (*iris.Application) {
    return app
}