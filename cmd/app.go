package cmd

import (
	"github.com/urfave/cli"
)

// 应用程序基本信息
const (
	AppName  = "Etcd Dump Tool"
	AppUsage = "dump etcd K/V to a file"
	Version  = "0.1.0"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = AppName
	app.Usage = AppUsage
	app.Version = Version
	app.Authors = []cli.Author{
		cli.Author{
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
