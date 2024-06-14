package utils

import (
	// "time"

	"github.com/spf13/viper"
)


type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	// TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	// AccessTokenDuration time.Duration `maptstructure:"ACCESS_TOKEN_DURATION"`
}


func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  //tells viper the path to look for a config
	viper.SetConfigName("app") //tells viper the name of the config file
	viper.SetConfigType("env")

	viper.AutomaticEnv() //reads matching env Variables from the environment if they exist

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}