package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type config struct {

	Server struct{
		Port string
	}

}

var Config config

func Init(path string) {

	index := strings.LastIndex(path, "/")

	p :=  path[0:index]
	c := path[index:]

	arr := strings.Split(c, ".")

	if len(arr) != 2 {
		log.Printf("configPath:%s, not find", path)
	}

	viper.SetConfigName(arr[0]) // name of config file (without extension)
	viper.SetConfigType(arr[1]) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(p)   // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Sprintf("unable to decode into struct, %v", err))
	}
}
