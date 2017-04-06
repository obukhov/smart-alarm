package interfaces

import (
	"time"
	"github.com/obukhov/smart-alarm/src/domain"
)

type AlarmActionRunner interface {
	CheckAndRun(timeToAlarm time.Duration, action domain.ActionInterface)
	Init(action domain.ActionInterface)
}

