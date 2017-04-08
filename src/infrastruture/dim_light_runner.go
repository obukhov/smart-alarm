package infrastruture

import (
	"time"
	"github.com/obukhov/smart-alarm/src/domain"
	"log"
	"errors"
	"os/exec"
	"github.com/dustin/go-humanize"
)

func NewDimLightRunner() *DimLightRunner {
	return &DimLightRunner{
		Config: struct {
			numberOfLevels int
			turnOn         string
			turnOff        string
			dimUp          string
			dimDown        string
			colorWhite     string
		}{
			numberOfLevels: 4,
			turnOn:         "LAMP_ON",
			turnOff:        "LAMP_OFF",
			dimUp:          "LAMP_DIM_UP",
			dimDown:        "LAMP_DIM_DOWN",
			colorWhite:     "LAMP_WHITE",
		},

	}
}

type DimLightRunner struct {
	Config struct {
		numberOfLevels int
		turnOn         string
		turnOff        string
		dimUp          string
		dimDown        string
		colorWhite     string
	}
	CurrentState struct {
		isOn     bool
		dimLevel uint8
	}
}

func (dlr *DimLightRunner) CheckAndRun(timeToAlarm time.Duration, action domain.ActionInterface) {
	dimLightAction, ok := action.(*domain.DimLightAction)
	if false == ok {
		panic(errors.New("Wrong action passed to DimLightRunner"))
	}

	timeToAlarmFormatter := humanize.RelTime(time.Now(), time.Now().Add(timeToAlarm), "earlier", "later")
	log.Printf("Check and run for %s", timeToAlarmFormatter)

	if timeToAlarm > dimLightAction.RumpUpDuration() {
		return
	}

	if dlr.CurrentState.isOn == false {
		dlr.send(dlr.Config.turnOn)
		dlr.CurrentState.isOn = true
	}

	step := int64(dimLightAction.RumpUpDuration()) / int64(dlr.Config.numberOfLevels)
	rumpDuration := int64(dimLightAction.RumpUpDuration() - timeToAlarm)
	desiredLevel := rumpDuration / step

	log.Printf("Rump duration is %d step is %d", rumpDuration, step)
	log.Printf("Desired level is %d current is %d", desiredLevel, dlr.CurrentState.dimLevel)

	if uint8(desiredLevel) > dlr.CurrentState.dimLevel {
		dlr.send(dlr.Config.dimUp)
		dlr.CurrentState.dimLevel++
	}
}

func (dlr *DimLightRunner) Init(action domain.ActionInterface) {
	dlr.send(dlr.Config.turnOn)
	dlr.send(dlr.Config.colorWhite)

	for i := 0; i < dlr.Config.numberOfLevels; i++ {
		dlr.send(dlr.Config.dimDown)
	}

	dlr.send(dlr.Config.turnOff)

}

func (dlr *DimLightRunner) send(commandName string) {
	log.Println("Sending commandName:" + commandName)
	command := exec.Command("irsend", "SEND_ONCE", "lamp", commandName)
	command.Start()
}
