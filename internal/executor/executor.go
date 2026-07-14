package executor

import (
	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/model"
)

func Apply(

	commands []model.Command,

	state *devices.State,

) {

	for _, cmd := range commands {

		switch cmd.Device {

		case "TargetOutput":

			state.TargetOutput = cmd.Action

		case "OperationalOutput":

			state.OperationalOutput = cmd.Action

		}

	}

}
