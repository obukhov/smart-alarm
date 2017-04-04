package domain

import (
	"gopkg.in/yaml.v2"
)

type ActionInterface interface {
	yaml.Marshaler
	yaml.Unmarshaler
	ActionType() string
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
