package bootstrap

import (
	"github.com/digimitzves/core/internal/config"
	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/engine"
	"github.com/digimitzves/core/internal/registry"
)

const (
	configPath = "data/config.json"
	eventsPath = "data/events.json"
	statePath  = "data/state.json"
)

func NewEngine() (*engine.Engine, error) {

	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, err
	}

	events, err := config.LoadEvents(eventsPath)
	if err != nil {
		return nil, err
	}

	state, err := devices.LoadState(statePath)
	if err != nil {
		return nil, err
	}

	reg := &registry.Registry{}

	e := &engine.Engine{
		Config:   cfg,
		Events:   events,
		Registry: reg,
		State:    state,
		Discover: devices.Discover,
	}

	return e, nil
}
