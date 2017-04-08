package usecase


type AlarmServiceCommand int

const (
	alarmServiceRefreshAlarm = iota
	alarmServiceStop
)

