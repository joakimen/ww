package main

import (
	"os"

	"github.com/joakimen/ww/cmd"
)

func main() {
	run(os.Args)
}

func run(args []string) {
	app := cmd.NewApp()
	err := app.Run(args)
	if err != nil {
		os.Exit(1)
	}
}
