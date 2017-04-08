package usecase

import (
	"github.com/obukhov/smart-alarm/src/domain"
	"github.com/obukhov/smart-alarm/src/interfaces"
	"time"
	"log"
	"errors"
)

type AlarmService struct {
	storage interfaces.AlarmStorage
	runners map[string]interfaces.AlarmActionRunner
	alarm   *domain.Alarm
	ticker  *time.Ticker
	command chan AlarmServiceCommand
}

func NewAlarmService(storage interfaces.AlarmStorage) *AlarmService {
	return &AlarmService{
		storage: storage,
		runners: make(map[string]interfaces.AlarmActionRunner),
	}
}

func (t *AlarmService) SetAlarm(alarm *domain.Alarm) {
	t.storage.Persist(alarm)
	t.alarm = alarm

	if nil != t.command {
		t.command <- alarmServiceRefreshAlarm
	}
}

func (t *AlarmService) ResetAlarm() {
	t.alarm = nil
}

func (t *AlarmService) LoadAlarm() {
	t.SetAlarm(t.storage.Load())
}

func (t *AlarmService) Start() {
	t.ticker = time.NewTicker(time.Second)
	t.command = make(chan AlarmServiceCommand)

	for _, action := range t.alarm.Actions {
		runner, ok := t.runners[action.ActionType()]
		if false == ok {
			panic(errors.New("Unknown action type " + action.ActionType()))
		}

		runner.Init(action)
	}

	go t.trackAlarm()
}

func (t *AlarmService) Stop() {
	t.ticker.Stop()
	t.ticker = nil
	close(t.command)
	t.command <- alarmServiceStop
}

func (t *AlarmService) AddRunner(alarmType string, runner interfaces.AlarmActionRunner) {
	t.runners[alarmType] = runner
}

func (t *AlarmService) trackAlarm() {
	alarm := *t.alarm //todo handle, when alarm is reset
	stop := false

	for stop == false {
		select {
		case now := <-t.ticker.C:
			for _, alarmAction := range alarm.Actions {
				actionType := alarmAction.ActionType()

				runner, found := t.runners[actionType]
				if false == found {
					log.Printf("Runner %s is not found for alarm", actionType)
					continue
				}

				runner.CheckAndRun(t.getTimeToAlarm(alarm, now), alarmAction)
			}
		case command := <-t.command:
			switch command {
			case alarmServiceRefreshAlarm:
				alarm = *t.alarm
			case alarmServiceStop:
				stop = true
			default:
				log.Printf("Unknown command: %d", command)

			}
		}
	}

	log.Println("Goroutine for alarm exited")
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
