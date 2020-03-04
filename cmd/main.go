package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/davidzaragoza/tangelo/pkg/data"
	"github.com/davidzaragoza/tangelo/pkg/domain"
	"github.com/davidzaragoza/tangelo/pkg/presentation"
)

func main() {
	confFile, err := os.Open("configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	config, err := readConfiguration(confFile)
	if err != nil {
		log.Fatal(err)
	}
	rep, err := data.NewRepository(config)
	if err != nil {
		log.Fatal(err)
	}
	uc := domain.NewUseCase(rep)
	server := presentation.NewServer(config, uc)
	server.StartServer()
}

func readConfiguration(file *os.File) (*domain.Configuration, error) {
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var result domain.Configuration
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
