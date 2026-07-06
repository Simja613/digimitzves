package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/digimitzves/core/internal/compiler"
	"github.com/digimitzves/core/internal/config"
	"github.com/digimitzves/core/internal/devices"
	"github.com/digimitzves/core/internal/scheduler"
	"github.com/digimitzves/core/internal/validator"
)

const configPath = "data/config.json"
const eventsPath = "data/events.json"
const statePath = "data/state.json"

func RegisterRoutes() {

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		w.Write([]byte(`{
			"service":"web",
			"status":"ok"
		}`))

	})

	http.HandleFunc("/api/config", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:

			cfg, err := config.Load(configPath)

			if err != nil {

				http.Error(w, err.Error(), 500)

				return

			}

			json.NewEncoder(w).Encode(cfg)

		case http.MethodPost:

			var cfg config.Config

			err := json.NewDecoder(r.Body).Decode(&cfg)

			if err != nil {

				http.Error(w, "invalid json", 400)

				return

			}

			err = validator.Validate(&cfg)

			if err != nil {

				http.Error(w, err.Error(), 400)

				return

			}

			err = config.Save(configPath, &cfg)

			if err != nil {

				http.Error(w, err.Error(), 500)

				return

			}

			json.NewEncoder(w).Encode(cfg)

		default:

			http.Error(
				w,
				"method not allowed",
				http.StatusMethodNotAllowed,
			)

		}

	})

	http.HandleFunc("/api/events", func(w http.ResponseWriter, r *http.Request) {

		events, err := config.LoadEvents(eventsPath)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		json.NewEncoder(w).Encode(events)

	})

	http.HandleFunc("/api/next-event", func(w http.ResponseWriter, r *http.Request) {

		events, err := config.LoadEvents(eventsPath)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		event, err := scheduler.FindNextEvent(
			events,
			time.Now(),
		)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		json.NewEncoder(w).Encode(event)

	})

	http.HandleFunc("/api/next-event-info", func(w http.ResponseWriter, r *http.Request) {

		events, err := config.LoadEvents(eventsPath)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		event, err := scheduler.FindNextEvent(
			events,
			time.Now(),
		)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		multi, err := scheduler.IsMultiDay(*event)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		start, _ := time.Parse(time.RFC3339, event.Start)

		end, _ := time.Parse(time.RFC3339, event.End)

		response := map[string]interface{}{

			"start": event.Start,

			"end": event.End,

			"duration_hours": end.Sub(start).Hours(),

			"multi_day": multi,
		}

		json.NewEncoder(w).Encode(response)

	})

	http.HandleFunc("/api/plan-preview", func(w http.ResponseWriter, r *http.Request) {

		events, err := config.LoadEvents(eventsPath)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		event, err := scheduler.FindNextEvent(
			events,
			time.Now(),
		)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		multi, err := scheduler.IsMultiDay(*event)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		cfg, err := config.Load(configPath)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		start, _ := time.Parse(time.RFC3339, event.Start)

		end, _ := time.Parse(time.RFC3339, event.End)

		response := map[string]interface{}{

			"event": event,

			"duration_hours": end.Sub(start).Hours(),

			"multi_day": multi,

			"behavior": cfg.MultiDayBehavior,
		}

		json.NewEncoder(w).Encode(response)

	})

	http.HandleFunc("/api/compile-preview", func(w http.ResponseWriter, r *http.Request) {

		events, err := config.LoadEvents(eventsPath)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		event, err := scheduler.FindNextEvent(
			events,
			time.Now(),
		)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		state, err := devices.LoadState(statePath)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		cfg, err := config.Load(configPath)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return

		}

		job := compiler.Compile(
			*event,
			*state,
			*cfg,
		)

		json.NewEncoder(w).Encode(job)

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("DigiMitzves Web OK"))

	})

}
