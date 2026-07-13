package engine

import (
	"time"

	"github.com/digimitzves/core/internal/config"
	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/model"
	"github.com/digimitzves/core/internal/registry"
)

func New(
	cfg *config.Config,
	job *model.Job,
	state *devices.State,
	reg *registry.Registry,
) *Engine {

	return &Engine{
		Config: cfg,

		Registry: reg,

		Job: job,

		State: state,

		Now: time.Now(),

		Context: Context{},
	}
}
