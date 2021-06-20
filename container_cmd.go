/**
 * @Author: hu5ky
 * @Description:
 * @File:  container
 * @Version: 0.0.1
 * @Date: 2020/12/6 8:35 下午
**/

package main

import (
	"dockerDemo/container"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: `create a container with namespace and cgroups limit dockerDemo run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpu share limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpu set limit",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := ctx.Args()
		tty := ctx.Bool("ti")
		//res := &subsystem.ResourceConfig{
		//	MemoryLimit: ctx.String("m"),
		//	CPUShare:    ctx.String("cpushare"),
		//	CPUSet:      ctx.String("cpuset"),
		//}
		container.Run(tty, cmd /*, res*/)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(ctx *cli.Context) error {
		logrus.Infof("Init come on")
		cmd := ctx.Args().Get(0)
		logrus.Infof("command %v", cmd)
		err := container.RunContainerInitProcess()
		return err
	},
}
