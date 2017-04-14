package domain

type Alarm struct {
	Enabled bool
	Hours   uint8 // todo make fields not published and provide serialization
	Minutes uint8

	Actions []ActionInterface
}

func NewAlarm(hours, minutes uint8, enabled bool, actions []ActionInterface) (*Alarm, typedError) {
	alarm := &Alarm{Enabled:enabled}
	if err := alarm.SetHours(hours); nil != err {
		return nil, err
	}

	if err := alarm.SetMinutes(minutes); nil != err {
		return nil, err
	}

	alarm.Actions = actions

	return alarm, nil
}

func (a *Alarm) SetHours(hours uint8) typedError {
	if hours < 0 || hours > 23 {
		return NewOutOfRangeErr(0, 23, int(hours))
	}

	a.Hours = hours

	return nil
}

func (a *Alarm) SetMinutes(minutes uint8) typedError {
	if minutes < 0 || minutes > 59 {
		return NewOutOfRangeErr(0, 59, int(minutes))
	}

	a.Minutes = minutes

	return nil
}
