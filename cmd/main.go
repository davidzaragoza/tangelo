package main

import (
	"github.com/davidzaragoza/tangelo/pkg/domain"
	"github.com/davidzaragoza/tangelo/pkg/presentation"
)

func main() {
	uc := domain.NewUseCase()
	server := presentation.NewServer(uc)
	server.StartServer()
}
