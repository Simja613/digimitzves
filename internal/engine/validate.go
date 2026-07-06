package engine

import "github.com/digimitzves/core/internal/model"

func (e *Engine) Validate() error {

	if e.Job == nil {
		return ErrInvalidJob
	}

	switch e.Job.Status {

	case model.JobPlanned,
		model.JobReady,
		model.JobRunning,
		model.JobCompleted,
		model.JobFailed,
		model.JobInvalid:

	default:
		return ErrInvalidJob

	}

	if e.Job.EventStart == "" {
		return ErrInvalidJob
	}

	if e.Job.EventEnd == "" {
		return ErrInvalidJob
	}

	if len(e.Job.Commands) == 0 {
		return ErrInvalidJob
	}

	for _, cmd := range e.Job.Commands {

		if cmd.Device == "" {
			return ErrInvalidJob
		}

		if cmd.Action == "" {
			return ErrInvalidJob
		}

		switch cmd.Status {

		case model.CommandPlanned,
			model.CommandRunning,
			model.CommandDone,
			model.CommandFailed:

		default:
			return ErrInvalidJob

		}

	}

	return nil

}
