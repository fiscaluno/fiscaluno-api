package handler

import (
    "fmt"
    "os"
    "github.com/urfave/cli"
    "github.com/fiscaluno/fiscaluno-api/database/migrations"
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
    command := c.Args().Get(0)
    if command == "migrate" {
        migrations.Migrate()
    } else if command == "teste" {
        fmt.Println("teste")
    }
    return nil
}
