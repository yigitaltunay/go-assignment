package config

import (
	"github.com/spf13/viper"
	"github.com/yigitaltunay/go-assignment/app/models"
	"github.com/yigitaltunay/go-assignment/app/services"
)

var BackendConfig models.Config
var ConfigFilePath = "config/config.json"

func ViperRead() {
	viper.SetConfigFile(ConfigFilePath) // optionally look for config in the working directory
	err := viper.ReadInConfig()                                                     // Find and read the config file
	if err != nil {
		panic(err)

	}
	BackendConfig.Providers = viper.GetStringSlice("providers")
	BackendConfig.Port = viper.GetString("backend.port")
}

func InitServe() {
	ViperRead()
	services.SetConfig(BackendConfig)

}
