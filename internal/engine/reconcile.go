package engine

import (
	"github.com/digimitzves/core/internal/compiler"
)

func (e *Engine) Reconcile() error {

	// нет активного события

	if e.Event == nil {

		e.Job = nil

		return nil

	}

	// Job еще не существует

	if e.Job == nil {

		job := compiler.Compile(
			*e.Event,
			*e.State,
			*e.Config,
		)

		e.Job = &job

		e.Context.SaveRequired = true

		return nil

	}

	// существующий Job пока считаем актуальным

	return nil

}
