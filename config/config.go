package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Nats struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Mqtt struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Service struct {
	NumOfWorker int  `mapstructure:"numOfWorker"`
	Debug       bool `mapstructure:"debug"`
	Nats        Nats `mapstructure:"nats"`
	Mqtt        Mqtt `mapstructure:"mqtt"`
}

func LoadFromFile(configPath string) (*Service, error) {
	viper.SetConfigType("json")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config %s. error: %w", configPath, err)
	}

	config := new(Service)
	err = viper.UnmarshalExact(&config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
