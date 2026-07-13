package engine

import (
	"fmt"

	"github.com/digimitzves/core/internal/registry"
)

func (e *Engine) Synchronize() error {

	discovered, err := e.Discover()
	if err != nil {
		return err
	}

	for _, d := range discovered {

		id := fmt.Sprintf("%s#%d", d.Parent, d.Channel)

		if e.Registry.Exists(id) {
			continue
		}

		e.Registry.Add(
			registry.Device{
				ID:           id,
				Parent:       d.Parent,
				FriendlyName: d.FriendlyName,
				Present:      true,
				Configured:   false,
				Ignored:      false,
			},
		)

		e.Context.NewDevices = true
		e.Context.ConfigurationRequired = true
		e.Context.SaveRequired = true
	}

	return nil
}
