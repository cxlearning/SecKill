package server

import (
	"github.com/urfave/cli"
	"myProject/SecKill/sk_proxy/setup"
)

var ProxyServer = cli.Command{
	Name:  "proxy_server",
	Usage: "sk_proxy server",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "config.toml",
			Usage: "toml配置文件",
		},
	},
	Action: proxyRun,
}

func proxyRun(c *cli.Context) {
	setup.Run(c.String("conf"))
}
