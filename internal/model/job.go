package model

type Job struct {
	Created string `json:"created"`

	Status JobStatus `json:"status"`

	EventStart string `json:"event_start"`

	EventEnd string `json:"event_end"`

	Commands []Command `json:"commands"`
}
