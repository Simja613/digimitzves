package devices

import (
	"encoding/json"
	"os"
)

type State struct {
	TargetOutput      string `json:"target_output"`
	OperationalOutput string `json:"operational_output"`

	TargetActive bool `json:"target_active"`
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
