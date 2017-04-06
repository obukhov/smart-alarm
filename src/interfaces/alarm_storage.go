package interfaces

import "github.com/obukhov/smart-alarm/src/domain"

type AlarmStorage interface {
	Persist(alarm *domain.Alarm)
	Load() *domain.Alarm
}
