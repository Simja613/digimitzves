package engine

import (
	"testing"

	"github.com/digimitzves/core/internal/config"
	"github.com/digimitzves/core/internal/devices"
)

func TestReconcileCreatesJob(t *testing.T) {

	e := &Engine{

		Config: &config.Config{

			Devices: map[string]config.DeviceConfig{

				"TargetOutput": {

					OperationalState: "OFF",

					TargetState: "ON",
				},
			},
		},

		State: &devices.State{

			TargetOutput: "OFF",

			TargetActive: false,
		},

		Event: &config.Event{

			Start: "2026-07-17T18:00:00Z",

			End: "2026-07-18T19:00:00Z",
		},
	}

	if e.Job != nil {

		t.Fatal("expected Job to be nil before Reconcile")

	}

	err := e.Reconcile()

	if err != nil {

		t.Fatal(err)

	}

	if e.Job == nil {

		t.Fatal("expected Job to be created")

	}

	if !e.Context.SaveRequired {

		t.Fatal("expected SaveRequired to be true")

	}

}
