package domain

import "time"

type State interface {
	TimeToAlarm() time.Duration
}

type ActionInterface interface {
	CheckAndRun(timeToAlarm time.Duration) State
	Init()
}

type Alarm struct {
	Hours   uint8
	Minutes uint8

	Actions []ActionInterface
}

func NewAlarm(hours, minutes uint8, actions []ActionInterface) (*Alarm, typedError) {
	if hours < 0 || hours > 23 {
		return nil, NewOutOfRangeErr(0, 23, int(hours))
	}

	if minutes < 0 || minutes > 59 {
		return nil, NewOutOfRangeErr(0, 59, int(minutes))
	}

	return &Alarm{hours, minutes, actions}, nil
}
