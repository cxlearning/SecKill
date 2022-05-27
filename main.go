package main

import (
	"github.com/urfave/cli"
	"myProject/SecKill/cmd/sk_admin/server"
	"myProject/SecKill/cmd/sk_admin/tool"

	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gate-super-backend"
	app.Commands = []cli.Command{
		server.AdminServer,
		tool.AdInit,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}


}



