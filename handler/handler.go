package handler

import (
    "os"
    "github.com/urfave/cli"
    "fmt"
    _"github.com/fiscaluno/fiscaluno-api/database/migrations"
    "go/importer"
)

var app *cli.App 

func Handle() {
    app = cli.NewApp()
    app.Name = "Fiscaluno"
    app.Usage = "Fiscaluno Command Handler"
    app.Action = handleCommand

    app.Run(os.Args)
}

func handleCommand(c *cli.Context) error {
    if c.Args().Get(0) == "migrate" {
        MigrateDatabase()
        return nil
    }
    return nil
}

func MigrateDatabase() {
    _, err := importer.Default().Import("github.com/fiscaluno/fiscaluno-api/database/migrations")
    if err != nil {
        fmt.Printf("error: %s\n", err.Error())
        return
    }
}
