package domain

import (
	"time"
	"encoding/json"
	"errors"
	"log"
	"github.com/dustin/go-humanize"
)

const (
	AlarmActionDimLight = "DimLightAction"
)

type ActionInterface interface {
	json.Marshaler
	ActionType() string
}

func NewDimLightAction(rumpUpDuration time.Duration) *DimLightAction {
	return &DimLightAction{rumpUpDuration: rumpUpDuration}
}

func NewDimLightActionFromMap(sourceMap map[string]interface{}) *DimLightAction {
	result := &DimLightAction{}

	rumpUpUntyped, ok := sourceMap["rumpUpDuration"]
	if false == ok {
		panic(errors.New("Failed reading actions"))
	}

	rumpUp, ok := rumpUpUntyped.(float64)
	if false == ok {
		panic(errors.New("Failed reading actions"))
	}
	result.rumpUpDuration = time.Duration(rumpUp)

	log.Printf("Set diration to %s", humanize.RelTime(time.Now(), time.Now().Add(result.rumpUpDuration), "", ""))

	return result
}

type DimLightAction struct {
	rumpUpDuration time.Duration
}

func (d *DimLightAction) ActionType() string {
	return AlarmActionDimLight
}

func (d *DimLightAction) RumpUpDuration() time.Duration {
	return d.rumpUpDuration
}

func (d *DimLightAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":           d.ActionType(),
		"rumpUpDuration": d.rumpUpDuration,
	})
}
