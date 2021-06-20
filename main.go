/**
 * @Author: hu5ky
 * @Description:
 * @File:  main
 * @Version: 0.0.1
 * @Date: 2020/12/1 9:37 下午
**/

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const (
	usage = `it is dockerDemo, for people who want to understand the underlying principle of docker,
		and it is a express edition.`

	name = "DockerDemo"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		logrus.SetFormatter(&logrus.JSONFormatter{})

		logrus.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

	//cgroups.Cgroups()
	//namespace.UTSNS()
	//namespace.IPCNS()
	//namespace.PIDNS()
	//namespace.MountNS()
	//namespace.USERNS()
	//namespace.NETNS()
}
