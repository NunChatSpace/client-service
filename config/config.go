package config

import (
	"log"

	"github.com/jinzhu/configor"
)

type Config struct {
	Credentials struct {
		SignString string `defaullt:"ThisIsSecretKey@ForJWTToken" env:"CREDENTIALS_SIGN_STRING"`
	}

	Services struct {
		IdentityService string `default:"http://backend.aistrom.com:8000" env:"SERVICES__IDENTITY_SERVICE"`
	}
}

func LoadConfig() (*Config, error) {
	var config Config
	var err error

	err = loadEnv()
	if err == nil {
		log.Println("Env file loaded")
	}

	err = configor.
		New(&configor.Config{AutoReload: false}).
		Load(&config)

	if err != nil {
		log.Println(err)
		log.Fatal("Error loading config")
	}

	return &config, err
}
