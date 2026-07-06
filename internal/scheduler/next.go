package scheduler

import (
	"time"

	"github.com/digimitzves/core/internal/config"
)

func FindNextEvent(events []config.Event, now time.Time) (*config.Event, error) {

	for _, event := range events {

		start, err := time.Parse(time.RFC3339, event.Start)

		if err != nil {

			return nil, err

		}

		if start.After(now) {

			return &event, nil

		}

	}

	return nil, nil

}
