package usecase

import "github.com/obukhov/smart-alarm/src/domain"

type TimerServiceInterface interface {
	SetAlarm(alarm *domain.Alarm)
}
