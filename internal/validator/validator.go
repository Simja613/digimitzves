package validator

import (
	"errors"

	"github.com/digimitzves/core/internal/config"
)

func Validate(cfg *config.Config) error {

	if cfg.MultiDayBehavior == "" {

		return errors.New("multi day behavior is required")

	}

	if cfg.MultiDayBehavior != "continuous" &&
		cfg.MultiDayBehavior != "daily_cycle" {

		return errors.New("unknown multi day behavior")

	}

	return nil
}
