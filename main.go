package main

import (

	"github.com/urfave/cli"
	"myProject/SecKill/cmd/sk_admin"

	"os"

)

func main() {
	app := cli.NewApp()
	app.Name = "gate-super-backend"
	app.Commands = []cli.Command{
		sk_admin.AdminServer,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}


}



