package usecase

import (
	"github.com/obukhov/smart-alarm/src/domain/api"
	"github.com/obukhov/smart-alarm/src/domain"
	"time"
)

func NewApiServer(alarmService *AlarmService) *ApiServer {
	return &ApiServer{alarmService: alarmService}
}

type ApiServer struct {
	alarmService *AlarmService
}

func (s *ApiServer) GetAlarm() api.AlarmGetCommandResult {
	alarm := s.alarmService.GetAlarm()

	return api.AlarmGetCommandResult{
		Alarms: []api.Alarm{
			{
				Hours:   alarm.Hours,
				Minutes: alarm.Minutes,
				Enabled: alarm.Enabled,
			},
		},
	}
}

func (s *ApiServer) SetAlarm(command api.AlarmSetRequestCommand) api.AlarmSetRequestResponse {
	if len(command.Alarms) == 0 {
		return api.AlarmSetRequestResponse{false, "Alarms are less then 1", }
	}

	alarm := command.Alarms[0]
	actions := []domain.ActionInterface{
		domain.NewDimLightAction(1 * time.Minute),
	}

	domainAlarm, err := domain.NewAlarm(
		alarm.Hours,
		alarm.Minutes,
		alarm.Enabled,
		actions,
	)

	if nil != err {
		return api.AlarmSetRequestResponse{false, err.Error()}
	}

	s.alarmService.SetAlarm(domainAlarm)

	return api.AlarmSetRequestResponse{true, ""}
}
