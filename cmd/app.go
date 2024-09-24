package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func NewApp() cli.App {
	return cli.App{
		Name:  "ww",
		Usage: "Creates tasks from data in various places",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Enable verbose output",
			},
		},
		Action: func(_ *cli.Context) error {
			fmt.Println("Hello, World!")
			return nil
		},
	}
}
