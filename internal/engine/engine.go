package engine

import (
	"time"

	"github.com/digimitzves/core/internal/config"
	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/model"
)

type Engine struct {
	Config *config.Config

	Events []config.Event

	Job *model.Job

	State *devices.State

	Now time.Time

	Context Context

	Mode SystemMode

	Event *config.Event
}
