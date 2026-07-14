package compiler

import (
	"time"

	"github.com/digimitzves/core/internal/config"
	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/model"
)

func Compile(
	event config.Event,
	state devices.State,
	cfg config.Config,
) model.Job {

	start, err := time.Parse(
		time.RFC3339,
		event.Start,
	)

	if err != nil {

		return model.Job{}

	}

	end, err := time.Parse(
		time.RFC3339,
		event.End,
	)

	if err != nil {

		return model.Job{}

	}

	commands := []model.Command{}

	// enter Target mode

	if !state.TargetActive {

		for device, settings := range cfg.Devices {

			current := getDeviceState(
				device,
				state,
			)

			if current != settings.TargetState {

				commands = append(
					commands,
					model.Command{
						Time:   start,
						Device: device,
						Action: settings.TargetState,
						Status: model.CommandPlanned,
					},
				)

			}

		}

	}

	// return to Operational mode

	for device, settings := range cfg.Devices {

		current := getDeviceState(
			device,
			state,
		)

		if current != settings.OperationalState {

			commands = append(
				commands,
				model.Command{
					Time:   end,
					Device: device,
					Action: settings.OperationalState,
					Status: model.CommandPlanned,
				},
			)

		}

	}

	job := model.Job{
		Created:    time.Now().Format(time.RFC3339),
		Status:     model.JobPlanned,
		EventStart: event.Start,
		EventEnd:   event.End,
		Commands:   commands,
	}

	return job

}

func getDeviceState(
	device string,
	state devices.State,
) string {

	switch device {

	case "TargetOutput":

		return state.TargetOutput

	case "OperationalOutput":

		return state.OperationalOutput

	default:

		return ""

	}

}
