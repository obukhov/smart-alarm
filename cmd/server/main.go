package main

import (
	"log"
	"github.com/obukhov/smart-alarm/src/infrastruture"
	"github.com/obukhov/smart-alarm/src/usecase"
	"github.com/obukhov/smart-alarm/src/domain"
)

func main() {
	log.Println("Starting server")

	storage := infrastruture.NewFileAlarmStorage()
	service := usecase.NewAlarmService(storage)

	service.AddRunner(domain.AlarmActionDimLight, infrastruture.NewDimLightRunner())

	service.LoadAlarm()
	service.Start()

	<-make(chan bool) // todo handle system signals for graceful stop
}
