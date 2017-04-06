package main

import (
	"log"
	"github.com/obukhov/smart-alarm/src/infrastruture"
	"github.com/obukhov/smart-alarm/src/usecase"
)

// build persister
// build service
// run
// block

func main() {
	log.Println("Starting server")

	storage := infrastruture.NewFileAlarmStorage()
	service := usecase.NewAlarmService(storage)

	service.Start()
}
