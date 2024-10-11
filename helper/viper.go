package helper

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
    vip := viper.New()
    vip.SetConfigName("config") // name of config file (without extension)
    vip.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
    vip.AddConfigPath(".")
    err := vip.ReadInConfig()
    if err != nil { // Handle errors reading the config file
    	panic(fmt.Errorf("fatal error config file: %w", err))
    }
    return vip
}
