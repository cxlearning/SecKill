package server

import (
	"github.com/urfave/cli"
	"myProject/SecKill/sk_layer/setup"
)

var LayerServer = cli.Command{
	Name:  "layer_server",
	Usage: "sk_layer server",
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
	Action: layerRun,
}

func layerRun(c *cli.Context) {
	setup.Run(c.String("conf"))
}
