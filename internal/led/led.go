package led

import (
	"log"

	"github.com/digimitzves/core/internal/health"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Show(status health.Status) {

	switch status {

	case health.OK:
		log.Println("LED: GREEN")

	case health.WARNING:
		log.Println("LED: YELLOW")

	case health.ERROR:
		log.Println("LED: RED")

	default:
		log.Println("LED: UNKNOWN")
	}
}
