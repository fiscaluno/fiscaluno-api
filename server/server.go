package server

import (
    "github.com/fiscaluno/fiscaluno-api/routes"
    "github.com/kataras/iris"
)

func Listen() {
    app := routes.Router()
    // Views Root Folder
    app.RegisterView(iris.HTML("./views", ".html").Reload(true))
    // On Error Code
    app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
        // .Values are used to communicate between handlers, middleware.
        errMessage := ctx.Values().GetString("error")
        if errMessage != "" {
            ctx.Writef("Internal server error: %s", errMessage)
            return
        }

        ctx.Writef("(Unexpected) internal server error")
    })

    app.Use(func(ctx iris.Context) {
        ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
        ctx.Next()
    })

    app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}

// Middleware Example
func logThisMiddleware(ctx iris.Context) {
    ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
    ctx.Next()
}
