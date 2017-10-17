package main

import (
	"github.com/urfave/cli"
	"path"
)

func BuildCmd(ex string, run func(string,string,string,string)) *cli.App {

	app := cli.NewApp()
	app.Name = path.Base(ex)
	app.Usage = "Mattermost team daily work checker"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "server",
			Usage:       "Mattermost server URL",
		},
		cli.StringFlag{
			Name:        "user",
			Usage:       "User login",
		},
		cli.StringFlag{
			Name:        "password",
			Usage:       "Password login",
		},
		cli.StringFlag{
			Name:        "channel",
			Usage:       "Mattermost channel to check",
		},
	}

	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		cli.ShowAppHelp(c)
		return err
	}

	app.Action = func(c *cli.Context) error {

		var flagsNum = c.NumFlags()
		if flagsNum != 4 {
			cli.ShowAppHelp(c)
			return cli.NewExitError("",ERROR_WRONG_ARGUMENTS_NUMBER)
		}

		var server = c.String("server")
		var user = c.String("user")
		var password = c.String("password")
		var channel = c.String("channel")

		if !CheckServer(server) {
			cli.ShowAppHelp(c)
			return cli.NewExitError("",ERROR_INVALID_SERVER_URL)
		}

		run(server, user, password, channel)
		return nil
	}

	return app
}
