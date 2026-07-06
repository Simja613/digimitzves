package devices

import (
	"encoding/json"
	"os"
)

type State struct {
	Channel1 string `json:"channel1"`
	Channel2 string `json:"channel2"`

	ShabbatMode bool `json:"shabbat_mode"`
}

func LoadState(path string) (*State, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var state State

	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func SaveState(path string, state *State) error {

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
