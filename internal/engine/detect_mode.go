package engine

func (e *Engine) DetectMode() error {

	e.Context.DetectedMode = NormalMode

	if e.Event != nil {
		e.Context.DetectedMode = EventMode
	}

	return nil

}
