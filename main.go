package main

import (
	"fmt"
	"os"

	"github.com/joakimen/ww/cmd"
	"github.com/joakimen/ww/pkg/credentials"
)

func main() {
	run(os.Args)
}

func run(args []string) {
	credentialsManager := credentials.NewKeychainCredentialsManager()
	app := cmd.NewApp(credentialsManager)
	err := app.Run(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
