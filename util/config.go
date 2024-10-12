package util

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config stores all configuration of tge application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
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

func FetchEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}

	return value
}
