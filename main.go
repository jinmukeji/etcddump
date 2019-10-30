package main

import (
	"log"
	"os"

	"github.com/jinmukeji/etcddump/cmd"
)

func main() {
	app := cmd.NewApp()
	err := app.Run(os.Args)
	die(err)
}

func die(err error) {
	if err != nil {
		log.Fatalf("ERROR: %v\r", err)
	}
}
