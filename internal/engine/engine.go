package engine

import (
	"time"

	"github.com/digimitzves/core/internal/config"
	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/model"
	"github.com/digimitzves/core/internal/registry"
)

type Engine struct {
	Config *config.Config

	Events []config.Event

	Registry *registry.Registry

	Job *model.Job

	State *devices.State

	Now time.Time

	Context Context

	Mode SystemMode

	Event *config.Event

	Discover func() ([]devices.DiscoveredDevice, error)
}
