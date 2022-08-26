package main

import (
	"github.com/tajpouria/goly/repository"
	"github.com/tajpouria/goly/server"
)

func main() {
	repository.Setup()
	server.SetupAndListen()
}
