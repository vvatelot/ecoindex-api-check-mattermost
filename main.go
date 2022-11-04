package main

import (
	"log"
	"os"
)

func main() {
	app := getCliApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
