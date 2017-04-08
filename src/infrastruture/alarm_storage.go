package infrastruture

import (
	"os"
	"github.com/obukhov/smart-alarm/src/domain"
	"encoding/json"
	"log"
	"io/ioutil"
	"errors"
	"fmt"
)

type UntypedHash map[string]interface{}

func NewFileAlarmStorage() *fileAlarmStorage {
	dir, err := os.Getwd()

	if nil != err {
		log.Println(err.Error())
		panic(err)
	}

	return &fileAlarmStorage{filePath: dir + "/alarm.json"}
}

type fileAlarmStorage struct {
	filePath string
}

func (as *fileAlarmStorage) Persist(alarm *domain.Alarm) {
	data, err := json.Marshal(alarm)
	if nil != err {
		panic(err)
	}

	err = ioutil.WriteFile(as.filePath, data, 0644)
	if nil != err {
		panic(err)
	}
}

func (as *fileAlarmStorage) Load() *domain.Alarm {
	alarm, typedErr := domain.NewAlarm(0, 0, make([]domain.ActionInterface, 0))
	if nil != typedErr {
		panic(typedErr)
	}

	data, err := ioutil.ReadFile(as.filePath)
	if nil != err {
		panic(err)
	}

	alarmRaw := make(UntypedHash)

	err = json.Unmarshal(data, &alarmRaw)
	if nil != err {
		panic(errors.New("Failed unmarshal json"))
	}

	log.Printf("JSON data is %s", string(data))

	alarm.SetHours(as.readUInt8(alarmRaw, "Hours"))
	alarm.SetMinutes(as.readUInt8(alarmRaw, "Minutes"))
	alarm.Actions = as.readActions(alarmRaw)

	log.Printf("Setting alarm %d:%d", alarm.Hours, alarm.Minutes)

	return alarm
}

func (as *fileAlarmStorage) readActions(alarmRaw UntypedHash) []domain.ActionInterface {
	result := make([]domain.ActionInterface, 0);
	actionsRawListUntyped, ok := alarmRaw["Actions"]
	if false == ok {
		panic(errors.New("Failed reading actions"))
	}

	actionsRawList, ok := actionsRawListUntyped.([]interface{}) //.([]map[string]interface{})
	if false == ok {
		panic(errors.New("Failed reading actions"))
	}

	for _, actionRawUntyped := range actionsRawList {
		actionRaw, ok := (actionRawUntyped).(map[string]interface{})
		if false == ok {
			panic(errors.New("Failed reading actions"))
		}

		result = append(result, as.readAction(actionRaw))
	}

	return result
}

func (as *fileAlarmStorage) readAction(actionRaw map[string]interface{}) domain.ActionInterface {
	actionTypeUntyped, ok := actionRaw["type"]
	if false == ok {
		panic(errors.New("Failed reading actions"))
	}

	actionType, ok := actionTypeUntyped.(string)
	if false == ok {
		panic(errors.New("Failed reading actions"))
	}

	switch actionType {
	case domain.AlarmActionDimLight:
		return domain.NewDimLightActionFromMap(actionRaw)
	default:
		panic(errors.New("Unknown action type"))
	}
}

func (as *fileAlarmStorage) readUInt8(alarmRaw map[string]interface{}, key string) uint8 {
	valueUntyped, ok := alarmRaw[key]
	if false == ok {
		panic(errors.New(fmt.Sprintf("Failed reading %s value", key)))
	}

	value, ok := valueUntyped.(float64)
	if false == ok {
		panic(errors.New(fmt.Sprintf("Failed reading %s value", key)))
	}

	return uint8(value)
}
