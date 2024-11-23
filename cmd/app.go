package cmd

import (
	"github.com/joakimen/ww/pkg/auth/aws"
	"github.com/joakimen/ww/pkg/auth/github"
	"github.com/joakimen/ww/pkg/auth/jira"
	"github.com/joakimen/ww/pkg/credentials"
	"github.com/urfave/cli/v2"
)

func NewApp(credentialsManager credentials.Manager) *cli.App {
	var (
		github = github.NewCredentialsManager(credentialsManager)
		jira   = jira.NewCredentialsManager(credentialsManager)
		aws    = aws.NewCredentialsManager(credentialsManager)
	)

	return &cli.App{
		Name:  "ww",
		Usage: "Creates tasks from data in various places",
		Commands: []*cli.Command{
			{
				Name:  "auth",
				Usage: "Manage authentication",
				Subcommands: []*cli.Command{
					{
						Name:  "show",
						Usage: "Show current authentication details",
						Action: func(_ *cli.Context) error {
							return credentialsManager.Show()
						},
					},
					{
						Name:  "delete",
						Usage: "Delete current authentication details",
						Action: func(_ *cli.Context) error {
							return credentialsManager.Delete()
						},
					},
				},
			},
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
			{
				Name:  "github",
				Usage: "Interact with GitHub",
				Subcommands: []*cli.Command{
					{
						Name:  "auth",
						Usage: "Manage GitHub authentication",
						Subcommands: []*cli.Command{
							{
								Name:  "login",
								Usage: "Login to GitHub",
								Action: func(_ *cli.Context) error {
									return github.Login()
								},
							},
							{
								Name:  "logout",
								Usage: "Logout from GitHub",
								Action: func(_ *cli.Context) error {
									return github.Logout()
								},
							},
							{
								Name:  "status",
								Usage: "Check authentication status",
								Action: func(_ *cli.Context) error {
									return github.Status()
								},
							},
							{
								Name:  "show",
								Usage: "Show current authentication details",
								Action: func(_ *cli.Context) error {
									return github.Show()
								},
							},
						},
					},
				},
			},
			{
				Name:  "aws",
				Usage: "Interact with AWS",
				Subcommands: []*cli.Command{
					{
						Name:  "auth",
						Usage: "Manage AWS authentication",
						Subcommands: []*cli.Command{
							{
								Name:  "login",
								Usage: "Login to AWS",
								Action: func(_ *cli.Context) error {
									return aws.Login()
								},
							},
							{
								Name:  "logout",
								Usage: "Logout from AWS",
								Action: func(_ *cli.Context) error {
									return aws.Logout()
								},
							},
							{
								Name:  "status",
								Usage: "Check authentication status",
								Action: func(_ *cli.Context) error {
									return aws.Status()
								},
							},
							{
								Name:  "show",
								Usage: "Show current authentication details",
								Action: func(_ *cli.Context) error {
									return aws.Show()
								},
							},
						},
					},
				},
			},
		},
	}
}
