package model

import "time"

type Command struct {
	Time   time.Time `json:"time"`
	Device string    `json:"device"`
	Action string    `json:"action"`

	Status     CommandStatus `json:"status"`
	ExecutedAt string        `json:"executed_at,omitempty"`
}
