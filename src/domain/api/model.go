package api

type Alarm struct {
	Hours   uint8
	Minutes uint8
	Enabled bool
}

type AlarmSetRequestCommand struct {
	Alarms []Alarm
}

type AlarmGetCommandResult struct {
	Alarms []Alarm
}

type AlarmSetRequestResponse struct {
	IsSuccessful bool
	ErrorMessage string `json:",omitempty"`
}
