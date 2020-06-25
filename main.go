package main

import (
	"context"
	"flag"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/mkuchenbecker/storage/service"
)

// Settings is the setting for the storage service.
type Settings struct {
	Port int `envconfig:"PORT" default:"9000"`
}

func getSettings(prefix string) *Settings {
	var s Settings
	err := envconfig.Process(prefix, &s)
	if err != nil {
		log.Fatal(context.Background(), err.Error())
	}
	return &s
}

func main() {
	flag.Parse()
	settings := getSettings("")
	service.New().Start(settings.Port)
}
