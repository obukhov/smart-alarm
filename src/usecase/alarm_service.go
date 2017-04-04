package usecase

import (
	"github.com/obukhov/smart-alarm/src/domain"
	"time"
	"github.com/coreos/fleet/log"
)

type AlarmStorage interface {
	Persist(alarm domain.Alarm)
	Load() domain.Alarm
}

type AlarmServiceCommand int

const (
	alarmServiceRefreshAlarm = iota
	alarmServiceStop
)

type AlarmActionRunner interface {
	CheckAndRun(timeToAlarm time.Duration, action domain.ActionInterface)
	Init(action domain.ActionInterface)
}

type AlarmService struct {
	storage AlarmStorage
	runners map[string]AlarmActionRunner
	alarm   domain.Alarm
	ticker  *time.Ticker
	command chan AlarmServiceCommand
}

func (t *AlarmService) SetAlarm(alarm domain.Alarm) {
	t.storage.Persist(alarm)
	t.alarm = alarm

	t.command <- alarmServiceRefreshAlarm
}

func (t *AlarmService) ResetAlarm() {
	t.ticker.Stop()
	t.command <- alarmServiceStop
}

func (t *AlarmService) Start() {
	t.alarm = t.storage.Load()
	t.ticker = time.NewTicker(time.Second)
	t.command = make(chan AlarmServiceCommand)

	go t.trackAlarm()

}

func (t *AlarmService) trackAlarm() {
	alarm := t.alarm
	stop := false

	for stop == false {
		select {
		case now := <-t.ticker.C:
			for _, alarmAction := range alarm.Actions {
				actionType := alarmAction.ActionType()

				runner, found := t.runners[actionType]
				if false == found {
					log.Errorf("Runner %s is not found for alarm", actionType)
					continue
				}

				runner.CheckAndRun(t.getTimeToAlarm(alarm, now), alarmAction)
			}
		case command := <-t.command:
			switch command {
			case alarmServiceRefreshAlarm:
				alarm = t.alarm
			case alarmServiceStop:
				stop = true
			default:
				log.Errorf("Unknown command: %d", command)

			}
		}
	}

	log.Info("Goroutine for alarm exited")
}

func (t *AlarmService) getTimeToAlarm(alarm domain.Alarm, currentTime time.Time) time.Duration {
	return t.getClosestAlarmTime(alarm, currentTime).Sub(currentTime);
}

func (t *AlarmService) getClosestAlarmTime(alarm domain.Alarm, targetDay time.Time) time.Time {
	if targetDay.Hour() > int(alarm.Hours) || (targetDay.Hour() == int(alarm.Hours) && targetDay.Minute() >= int(alarm.Minutes)) {
		targetDay = targetDay.Add(24 * time.Hour)
	}

	alarmTime := time.Date(
		targetDay.Year(), targetDay.Month(), targetDay.Day(),
		int(alarm.Hours), int(alarm.Minutes), 0, 0,
		targetDay.Location(),
	)

	return alarmTime
}
