package main

import (
	"fmt"
	"regexp"
)

func main() {
	//app := cli.NewApp()
	//app.Name = "gate-super-backend"
	//app.Commands = []cli.Command{
	//	server.AdminServer,
	//	server.LayerServer,
	//	tool.AdminInit,
	//}
	//
	//err := app.Run(os.Args)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(checkTelePhone("15869397390"))

}

func checkTelePhone(nums string) bool {
	regular := "^[1][3,4,5,6,7,8,9][0-9]{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(nums)
}

