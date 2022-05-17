package utils

import (
	"time"

	"github.com/spf13/viper"
)

//stores all configurations for the app read by viper either from file or env
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSourceSting       string        `mapstructure:"DB_SOURCE_STRING"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
