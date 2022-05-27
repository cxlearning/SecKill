package main

import (
	"github.com/urfave/cli"
	"myProject/SecKill/cmd/server"
	"myProject/SecKill/cmd/tool"

	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gate-super-backend"
	app.Commands = []cli.Command{
		server.AdminServer,
		tool.AdminInit,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}


}



