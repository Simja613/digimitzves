package engine

import (
	"time"

	"github.com/digimitzves/core/internal/config"
	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/model"
)

func New(
	cfg *config.Config,
	job *model.Job,
	state *devices.State,
) *Engine {

	return &Engine{
		Config:  cfg,
		Job:     job,
		State:   state,
		Now:     time.Now(),
		Context: Context{},
	}

}
