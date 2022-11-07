package main

import (
	"fmt"

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
	ecoindexUrl := cCtx.String("ecoindex_url")
	mattermostUrl := cCtx.String("mattermost_url")
	name := cCtx.String("name")

	setLockFileName(name, ecoindexUrl)
	errEcoindex := getEcoindexHealth(ecoindexUrl)

	if errEcoindex != nil && !isLockFileExist() {
		message := fmt.Sprintf("🚨 **ecoindex API error** \n\nThe Ecoindex API %s is in error. Here is the detail: \n```%s```", name, errEcoindex.Error())
		createLockFile()
		sendMessage(message, mattermostUrl)
	}

	if errEcoindex == nil && isLockFileExist() {
		message := fmt.Sprintf("✅ **ecoindex API OK** \n\nThe Ecoindex API %s is now OK.", name)
		removeLockFile()
		sendMessage(message, mattermostUrl)
	}

	return nil
}
