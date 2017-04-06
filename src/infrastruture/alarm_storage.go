package infrastruture

import (
	"os"
	"github.com/obukhov/smart-alarm/src/domain"
	"encoding/json"
	"log"
	"io/ioutil"
)

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

	json.Unmarshal(data, alarm)

	return alarm
}
