package scheduler

import (
	"time"

	"github.com/digimitzves/core/internal/config"
)

func FindCurrentEvent(
	events []config.Event,
	now time.Time,
) (*config.Event, error) {

	for i := range events {

		start, err := time.Parse(
			time.RFC3339,
			events[i].Start,
		)

		if err != nil {
			return nil, err
		}

		end, err := time.Parse(
			time.RFC3339,
			events[i].End,
		)

		if err != nil {
			return nil, err
		}

		if !now.Before(start) && now.Before(end) {

			return &events[i], nil

		}

	}

	return nil, nil

}
