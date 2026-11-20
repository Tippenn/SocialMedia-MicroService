package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server         *Server
		AccountServer  *Server
		TimelineServer *Server
		PostServer     *Server
		Redis          *Redis
		Auth           *Auth
	}

	Server struct {
		Port    int
		BaseURL string
	}

	Redis struct {
		Host     string
		Port     int
		Password string
	}
	Auth struct {
		Secret string
	}
)

var (
	once   sync.Once
	config *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		err := viper.Unmarshal(&config)
		if err != nil {
			panic(err)
		}
	})

	return config
}
