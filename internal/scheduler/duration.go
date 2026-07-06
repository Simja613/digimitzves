package scheduler

import (
	"time"

	"github.com/digimitzves/core/internal/config"
)

func IsMultiDay(event config.Event) (bool, error) {

	start, err := time.Parse(
		time.RFC3339,
		event.Start,
	)

	if err != nil {

		return false, err

	}

	end, err := time.Parse(
		time.RFC3339,
		event.End,
	)

	if err != nil {

		return false, err

	}

	duration := end.Sub(start)

	return duration > 30*time.Hour, nil

}
