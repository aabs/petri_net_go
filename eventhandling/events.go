package eventhandling

import (
	"aabs/petri_net_go/core"
	"time"

	"github.com/google/uuid"
)

func CreateEvent(version, eventType string, payload interface{}) (Event, error) {
	eventId := uuid.New().String()
	return Event{
		EventId:            eventId,
		Version:            version,
		EventType:          eventType,
		GeneratedTimestamp: time.Now().UnixNano(),
		Payload:            payload,
	}, nil
}

type EventHandler interface {
	HandleEvent(*core.Marking, Event) (*core.Marking, error)
}

type StateInjector interface {
	InjectState(*core.Marking, Event) (*core.Marking, error)
}

func SetMarking(m *core.Marking, placeId int, token float64) (*core.Marking, error) {
	return m.SetMarking(placeId, token)
}
