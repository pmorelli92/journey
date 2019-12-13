package main

import (
	"log"
	"os"

	journey "github.com/pmorelli92/journey/commands"
	_ "github.com/db-journey/postgresql-driver"
	"github.com/urfave/cli"
)

func main() {
	app := App()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func App() *cli.App {
	app := cli.NewApp()
	app.Usage = "Migrations and cronjobs for databases (postgres)"
	app.Version = "2.4.0"

	app.Flags = journey.Flags()

	app.Commands = journey.Commands()
	return app
}
