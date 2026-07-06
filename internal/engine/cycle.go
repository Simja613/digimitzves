package engine

import "time"

func (e *Engine) Cycle() error {

	e.Now = time.Now()

	// ---------- OBSERVE ----------

	if err := e.Validate(); err != nil {
		return err
	}

	if err := e.DetectEvent(); err != nil {
		return err
	}

	if err := e.DetectMode(); err != nil {
		return err
	}

	// ---------- DECIDE ----------

	if err := e.Recover(); err != nil {
		return err
	}

	if err := e.Synchronize(); err != nil {
		return err
	}

	if err := e.Reconcile(); err != nil {
		return err
	}

	// ---------- ACT ----------

	if err := e.Execute(); err != nil {
		return err
	}

	// ---------- PERSIST ----------

	if err := e.Finalize(); err != nil {
		return err
	}

	return nil

}
