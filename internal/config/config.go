package config

import (
	"encoding/json"
	"os"
)

type DeviceConfig struct {
	OperationalState string `json:"operational_state"`

	TargetState string `json:"target_state"`
}

type Config struct {
	MultiDayBehavior string `json:"multi_day_behavior"`

	Devices map[string]DeviceConfig `json:"devices"`
}

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)

	if err != nil {

		return nil, err

	}

	var cfg Config

	err = json.Unmarshal(data, &cfg)

	if err != nil {

		return nil, err

	}

	return &cfg, nil

}

func Save(path string, cfg *Config) error {

	data, err := json.MarshalIndent(
		cfg,
		"",
		"  ",
	)

	if err != nil {

		return err

	}

	return os.WriteFile(
		path,
		data,
		0644,
	)

}
