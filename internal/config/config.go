package config

import (
	"github.com/spf13/viper"
	"log"
)

// Read attempts to read a config file from path, defaults to provided values if no config file is present
func Read(path string, defaults map[string]interface{}) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Println("proceeding with defaults")
		for key, val := range defaults {
			v.SetDefault(key, val)
		}
	}
	return v
}
