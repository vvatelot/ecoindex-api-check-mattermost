package main

import (
	"github.com/urfave/cli/v2"
)

func getCliApp() *cli.App {
	return &cli.App{
		Name:  "Ecoindex Bot",
		Usage: "Check Ecoindex API Status and alert in case of emergency",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ecoindex_url",
				Value: "http://localhost:8001",
				Usage: "base url of the ecoindex API to check",
			},
			&cli.StringFlag{
				Name:  "mattermost_url",
				Usage: "mattermost webhook url used to send message",
			},
			&cli.StringFlag{
				Name:  "name",
				Value: "Local",
				Usage: "name of the ecoindex API. It can bes something like 'production', 'backup'",
			},
		},
		Action: checkEcoindex,
	}
}

func checkEcoindex(cCtx *cli.Context) error {
	url := cCtx.String("ecoindex_url")
	mattermostUrl := cCtx.String("mattermost_url")
	name := cCtx.String("name")

	if err := isEcoindexHealthy(url); err != nil {
		if err := sendMessage(err.Error(), mattermostUrl, name); err != nil {
			return err
		}

	}

	return nil
}
