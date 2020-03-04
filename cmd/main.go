package main

import (
	"github.com/davidzaragoza/tangelo/pkg/data"
	"github.com/davidzaragoza/tangelo/pkg/domain"
	"github.com/davidzaragoza/tangelo/pkg/presentation"
)

func main() {
	rep := data.NewRepository()
	uc := domain.NewUseCase(rep)
	server := presentation.NewServer(uc)
	server.StartServer()
}
