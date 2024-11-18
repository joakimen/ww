package cmd

import (
	"github.com/joakimen/ww/internal/jira"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:  "ww",
		Usage: "Creates tasks from data in various places",
		Commands: []*cli.Command{
			{
				Name:  "jira",
				Usage: "Interact with Jira",
				Subcommands: []*cli.Command{
					{
						Name:  "auth",
						Usage: "Manage Jira authentication",
						Subcommands: []*cli.Command{
							{
								Name:  "login",
								Usage: "Login to Jira",
								Action: func(_ *cli.Context) error {
									return jira.Login()
								},
							},
							{
								Name:  "logout",
								Usage: "Logout from Jira",
								Action: func(_ *cli.Context) error {
									return jira.Logout()
								},
							},
							{
								Name:  "status",
								Usage: "Check authentication status",
								Action: func(_ *cli.Context) error {
									return jira.Status()
								},
							},
							{
								Name:  "show",
								Usage: "Show current authentication details",
								Action: func(_ *cli.Context) error {
									return jira.Show()
								},
							},
						},
					},
				},
			},
		},
	}
}
