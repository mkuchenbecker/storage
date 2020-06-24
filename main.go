package main

import (
	"github.com/mkuchenbecker/storage/service"
)

const (
	port = 50051
)

func main() {
	service.New().Start(port)
}
