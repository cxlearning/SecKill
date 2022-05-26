package sk_admin

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
		cli.StringFlag{
			Name:  "args",
			Value: "",
			Usage: "multiconfig cmd line args",
		},
	},
	Action: run,
}

func run(c *cli.Context) {
	setup.Run(c.String("conf"))
}




