package infrastruture

import (
	"os"
	"github.com/coreos/fleet/log"
)

func NewFileAlarmStorage() *fileAlarmStorage {
	dir, err := os.Getwd()

	if nil != err {
		log.Error(err.Error())
		panic(err)
	}

	return &fileAlarmStorage{filePath: dir}
}

type fileAlarmStorage struct {
	filePath string
}
