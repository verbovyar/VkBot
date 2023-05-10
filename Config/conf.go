package Config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ApiKey string `mapstructure:"API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("conf")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("do not parse config file:%v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		_ = fmt.Errorf("do not parse config file:%v", err)
	}

	return
}
