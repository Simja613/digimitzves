package engine

import "github.com/digimitzves/core/internal/scheduler"

func (e *Engine) DetectEvent() error {

	event, err := scheduler.FindCurrentEvent(
		e.Events,
		e.Now,
	)

	if err != nil {
		return err
	}

	e.Event = event

	return nil

}
