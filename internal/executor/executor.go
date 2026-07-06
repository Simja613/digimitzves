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

		case "channel1":

			state.Channel1 = cmd.Action

		case "channel2":

			state.Channel2 = cmd.Action

		}

	}

}
