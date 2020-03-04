package main

import (
	"github.com/davidzaragoza/tangelo/pkg/presentation"
)

func main() {
	server := presentation.NewServer()
	server.StartServer()
}
