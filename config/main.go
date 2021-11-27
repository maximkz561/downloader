package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type App struct {
	Port     string `default:"8000"`
	LogLevel string
}

type Mongo struct {
	Host   string `default:"127.0.0.1"`
	Port   string `default:"27017"`
	DbName string `default:"downloader"`
}

type Config struct {
	App   App
	Mongo Mongo
}

func GetConfig() Config {
	var app App
	var mongo Mongo
	err := envconfig.Process("app", &app)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("app", &mongo)
	if err != nil {
		log.Fatal(err.Error())
	}
	return Config{
		App:   app,
		Mongo: mongo,
	}
}
