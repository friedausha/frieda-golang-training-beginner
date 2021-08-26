package util

import (
	_ "context"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func MustHaveEnv(key string) string {
	env := viper.GetString(key)
	if env == "" {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err, "can't read .env file")
		}
		env = viper.GetString(key)
	}

	return env
}
