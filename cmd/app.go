package cmd

import (
	"github.com/urfave/cli"
)

const (
	appName  = "Etcd Dump Tool"
	appUsage = "dump etcd K/V to a file"
	version  = "0.1.0"
)

// NewApp creates a new CLI app.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = appName
	app.Usage = appUsage
	app.Version = version
	app.Authors = []cli.Author{
		{
			Name:  "Sky Jia",
			Email: "me@skyjia.com",
		},
	}

	app.Commands = []cli.Command{
		dumpCmd(),
		restoreCmd(),
	}

	return app
}
