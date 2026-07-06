package engine

import "time"

func (e *Engine) Run() {

	for {

		_ = e.Cycle()

		time.Sleep(time.Second)

	}

}
