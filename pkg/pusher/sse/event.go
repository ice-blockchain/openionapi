package sse

import (
	"github.com/ice-blockchain/openionapi/pkg/pusher/events"
)

type Event struct {
	Name    events.Name
	EventID int64  `json:"event_id"`
	Data    []byte `json:"data"`
}
