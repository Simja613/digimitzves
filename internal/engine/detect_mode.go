package engine

func (e *Engine) DetectMode() error {

	e.Context.DetectedMode = OperationalMode

	if e.Event != nil {
		e.Context.DetectedMode = TargetMode
	}

	return nil

}
