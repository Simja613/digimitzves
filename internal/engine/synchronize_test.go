package engine

import (
	"testing"

	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/registry"
)

func TestSynchronizeNewDevice(t *testing.T) {

	e := &Engine{

		Registry: &registry.Registry{},

		Discover: func() ([]devices.DiscoveredDevice, error) {

			return []devices.DiscoveredDevice{

				{
					Parent:       "relay",
					Channel:      1,
					FriendlyName: "Kitchen",
				},
			}, nil

		},
	}

	err := e.Synchronize()

	if err != nil {
		t.Fatal(err)
	}

	if len(e.Registry.Devices) != 1 {
		t.Fatal("device was not registered")
	}

	if !e.Context.NewDevices {
		t.Fatal("expected NewDevices flag")
	}

	if !e.Context.ConfigurationRequired {
		t.Fatal("expected ConfigurationRequired flag")
	}

	if !e.Context.SaveRequired {
		t.Fatal("expected SaveRequired flag")
	}

}
