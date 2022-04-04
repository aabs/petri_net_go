package persistence

import (
	"aabs/petri_net_go/core"
	"errors"
	"time"

	"github.com/google/uuid"
	"gonum.org/v1/gonum/mat"
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
	if placeId < 0 || placeId >= m.Places.Len() {
		return nil, errors.New("placeId out of range")
	}
	m_1 := mat.VecDenseCopyOf(m.Places)
	m_1.SetVec(placeId, token)
	return &core.Marking{Places: m_1}, nil
}
