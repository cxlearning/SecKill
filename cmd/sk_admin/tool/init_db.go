package tool

import (
	"github.com/urfave/cli"
	"myProject/SecKill/sk_admin/setup"
)

var AdInit = cli.Command{
	Name:  "ad_init",
	Usage: "init db",
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
	setup.InitDB(c.String("conf"))
}
