package utils

type Config struct {
	Version     string `mapstructure:"version"`
	AppName     string `mapstructure:"app_name"`
	Environment string `mapstructure:"environment"`
}
