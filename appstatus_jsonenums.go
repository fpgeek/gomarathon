// generated by jsonenums -type=AppStatus; DO NOT EDIT

package gomarathon

import (
	"encoding/json"
	"fmt"
)

var (
	_AppStatusNameToValue = map[string]AppStatus{
		"None":      AppStatusNone,
		"Healthy":   AppStatusHealthy,
		"UnHealthy": AppStatusUnHealthy,
		"Scaling":   AppStatusScaling,
		"Running":   AppStatusRunning,
	}

	_AppStatusValueToName = map[AppStatus]string{
		AppStatusNone:      "None",
		AppStatusHealthy:   "Healthy",
		AppStatusUnHealthy: "UnHealthy",
		AppStatusScaling:   "Scaling",
		AppStatusRunning:   "Running",
	}
)

func init() {
	var v AppStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_AppStatusNameToValue = map[string]AppStatus{
			interface{}(AppStatusNone).(fmt.Stringer).String():      AppStatusNone,
			interface{}(AppStatusHealthy).(fmt.Stringer).String():   AppStatusHealthy,
			interface{}(AppStatusUnHealthy).(fmt.Stringer).String(): AppStatusUnHealthy,
			interface{}(AppStatusScaling).(fmt.Stringer).String():   AppStatusScaling,
			interface{}(AppStatusRunning).(fmt.Stringer).String():   AppStatusRunning,
		}
	}
}

func (r AppStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _AppStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid AppStatus: %d", r)
	}
	return json.Marshal(s)
}

func (r *AppStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AppStatus should be a string, got %s", data)
	}
	v, ok := _AppStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid AppStatus %q", s)
	}
	*r = v
	return nil
}
