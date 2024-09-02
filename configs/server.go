package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Server struct {
	Port int `envconfig:"HTTP_PORT" default:"8080"`
}

func InitServerConfig() Server {
	var cfg Server
	err := godotenv.Overload()
	if err != nil {
		log.Println(err)
	}
	envconfig.MustProcess("", &cfg)
	return cfg
}
