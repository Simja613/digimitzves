package config

import (
	"encoding/json"
	"os"
)

type Event struct {
	Start string `json:"start"`

	End string `json:"end"`
}

func LoadEvents(path string) ([]Event, error) {

	data, err := os.ReadFile(path)

	if err != nil {

		return nil, err

	}

	var events []Event

	err = json.Unmarshal(data, &events)

	if err != nil {

		return nil, err

	}

	return events, nil

}
