package event

import (
	"utils"
	"github.com/go-clean-archituecture/valueobjects"
)
type Event struct {
	aggregateId string
	aggregateType string
	version int
	reason string
	timestamp valueobjects.Timestamp
	payload event.EventPayload
	observerList []Observer
	serializerType  valueobjects.SerializerType
}

func(e *Event) GetAggregateId() string {
	return e.aggregateId
}

func(e *Event) SetAggregateId(id string) {
	return e.aggregateId  = id
}

func(e *Event) GetAggregateType() string {
	return e.aggregateType
}

func(e *Event) SetAggregateType(aggType string) {
	return e.aggregateType  = aggType
}

func(e *Event) GetVersion() int {
	return e.version
}

func(e *Event) SetVersion(version int) {
	return e.version  = version
}

func(e *Event) GetReason() string {
	return e.reason
}

func(e *Event) SetReason(aggReason string) {
	return e.reason  = aggReason
}

func(e *Event) GetTimestamp() valueobjects.Timestamp {
	return e.timestamp
}

func(e *Event) SetTimestamp(timestamp valueobjects.Timestamp) {
	return e.timestamp  = timestamp
}

func(e *Event) GetPayload() valueobjects.EventPayload {
	return e.payload
}

func(e *Event) SetPayload(payload valueobjects.EventPayload) {
	return e.payload  = payload
}

func (evt *Event) GetSerializer() valueobjects.SerializerType {
	return evt.serializerType
}

func (evt *Event) SetSerializer(typ valueobjects.SerializerType) {
	evt.serializerType = typ
}

func (i *Event) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Event) Deregister(o Observer) {
	i.observerList = utils.RemoveFromslice(i.observerList, o)
}

func (i *Event) NotifyAll() {
	for _, observer := range i.observerList {
			observer.Excute(i.name)
	}
}
