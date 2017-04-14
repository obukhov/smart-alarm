package infrastruture

import (
	"github.com/obukhov/smart-alarm/src/usecase"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"github.com/obukhov/smart-alarm/src/domain/api"
)

func NewApiServer(alarmServer *usecase.AlarmService) *rest.Api {
	apiServer := usecase.NewApiServer(alarmServer)

	apiApp := rest.NewApi()
	apiApp.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/alarms", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(apiServer.GetAlarm())
		}),

		rest.Post("/alarms", func(w rest.ResponseWriter, req *rest.Request) {
			requestCommand := api.AlarmSetRequestCommand{}
			err := req.DecodeJsonPayload(&requestCommand)
			if nil != err {
				w.WriteJson(api.AlarmSetRequestResponse{
					false,
					"Wrong json format",
				})
			}

			resp := apiServer.SetAlarm(requestCommand)
			w.WriteJson(resp)
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	apiApp.SetApp(router)

	return apiApp
}
