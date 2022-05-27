package server

import (
	"github.com/urfave/cli"
	"myProject/SecKill/sk_admin/setup"
)
var AdminServer = cli.Command{
	Name:  "ad_server",
	Usage: "sk_admin server",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "config.toml",
			Usage: "toml配置文件",
		},
	},
	Action: run,
}

func run(c *cli.Context) {
	setup.Run(c.String("conf"))
}




