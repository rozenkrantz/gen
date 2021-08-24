package main

import (
	"github.com/dequinox/gen/generate"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	mainInfo    = "General info"
	projectInfo = "Project name"
)

var initFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    mainInfo,
		Aliases: []string{"m"},
		Value:   "main.go",
		Usage:   "Main go file",
	},
	&cli.StringFlag{
		Name:    projectInfo,
		Aliases: []string{"p"},
		Value:   "my-project",
		Usage:   "Name of the project",
	},
}

func initAction(c *cli.Context) error {
	return generate.Build(&generate.Config{
		MainFile: c.String(mainInfo),
	})
}

func main() {
	app := cli.NewApp()
	app.Usage = "Automatically generate project skeleton"
	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create skeleton",
			Action:  initAction,
			Flags:   initFlags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
