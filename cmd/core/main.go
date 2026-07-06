package main

import (
	"log"

	"github.com/digimitzves/core/internal/health"
	"github.com/digimitzves/core/internal/led"
	"github.com/digimitzves/core/internal/scheduler"
	timeservice "github.com/digimitzves/core/internal/time"
)

func main() {

	log.Println("DigiMitzves Core starting")

	timeService := timeservice.New()
	schedulerService := scheduler.New()
	healthService := health.New()
	ledService := led.New()

	_ = timeService
	_ = schedulerService

	status := healthService.Overall()

	log.Println("DigiMitzves Core initialized")
	log.Printf("System health status: %d", status)

	ledService.Show(status)
}
