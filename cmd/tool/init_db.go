package tool

import (
	"github.com/urfave/cli"
	"myProject/SecKill/sk_admin/setup"
)

var AdminInit = cli.Command{
	Name:  "ad_init",
	Usage: "init db",
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
	setup.InitDB(c.String("conf"))
}
