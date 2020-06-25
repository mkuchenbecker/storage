package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/golang/glog"
	"github.com/kelseyhightower/envconfig"
	"github.com/mkuchenbecker/storage/api"
	"github.com/mkuchenbecker/storage/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Settings is the setting for the storage service.
type Settings struct {
	Port int `envconfig:"PORT" default:"8080"`
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
	storage := service.New()

	glog.Infof("Starting Service on Port: %d", settings.Port)
	defer glog.Flush()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", settings.Port))
	if err != nil {
		panic(errors.Wrap(err, "failed to listen"))
	}
	grpcServer := grpc.NewServer()
	api.RegisterStorageServer(grpcServer, storage)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
