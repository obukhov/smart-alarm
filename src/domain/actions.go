package domain

import (
	"time"
	"encoding/json"
)

type ActionInterface interface {
	json.Marshaler
	ActionType() string
}

func NewDimLightAction(rumpUpDuration time.Duration) *DimLightAction {
	return &DimLightAction{rumpUpDuration: rumpUpDuration}
}

type DimLightAction struct {
	rumpUpDuration time.Duration
}

func (d *DimLightAction) ActionType() string {
	return "DimLightAction"
}

func (d *DimLightAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":           d.ActionType(),
		"rumpUpDuration": d.rumpUpDuration,
	})
}
