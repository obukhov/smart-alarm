package main

import (
	"log"
	"github.com/obukhov/smart-alarm/src/infrastruture"
	"github.com/obukhov/smart-alarm/src/usecase"
	"github.com/obukhov/smart-alarm/src/domain"
	"net/http"
)

func main() {
	log.Println("Starting server")

	storage := infrastruture.NewFileAlarmStorage()
	service := usecase.NewAlarmService(storage)

	service.AddRunner(domain.AlarmActionDimLight, infrastruture.NewDimLightRunner())
	service.LoadAlarm()
	service.Start()


	apiServer := infrastruture.NewApiServer(service)
	log.Fatal(http.ListenAndServe(":8090", apiServer.MakeHandler()))
}
