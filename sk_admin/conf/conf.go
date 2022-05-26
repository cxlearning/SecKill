package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	Port int
	Name string
}

var Config config

func Init(path string) {
	fmt.Println("path====", path)

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)   // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Sprintf("unable to decode into struct, %v", err))
	}
}
