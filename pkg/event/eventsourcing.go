package event

import (
	"errors"
	"reflect"
)

type AggregateCluster struct {
	currentId string
	currentType string
	currentVersion int
	committedEvent []Event
	uncommitedEvent []Event
	transitionfn	Transition
}

func(agc *AggregateCluster) GetId() string {
	return agc.currentId
}

func(agc *AggregateCluster) SetId(id string) {
	return agc.currentId  = id
}

func(agc *AggregateCluster) GetType() string {
	return agc.currentType
}

func(agc *AggregateCluster) SetType(aggType string) {
	return agc.currentType  = aggType
}

func(agc *AggregateCluster) GetVersion() int {
	return agc.currentVersion
}

func(e *AggregateCluster) SetVersion(version int) {
	return agc.currentVersion  = version
}

func(agc *AggregateCluster) Apply(event Event) error {
	return agc.apply(event, false)
}
// ApplyCommitted applies already committed event. The AggregateCluster state
// id, type, version will be replaced with current event id, type and version.
func (r *AggregateCluster) ApplyCommitted(evt event.Eventer) error {
	return r.apply(evt, true)
}

func(agc *AggregateCluster) apply(event Event, commited bool) error {
	if err := agc.transitionfn(event); err != nil {
		return err
	}
	if committed {
		if err := r.checkVersionDuplication(evt); err != nil {
			return err
		}

		r.currentId = evt.GetAggregateId()
		r.currentType = evt.GetAggregateType()
		r.currentVersion = evt.GetVersion()
		r.committedEvents = append(r.committedEvents, evt)
	} else {
		// Increment our aggregate root version for +1
		r.currentVersion = r.nextVersion()

		evt.SetAggregateId(r.currentId)
		evt.SetAggregateType(r.currentType)
		evt.SetVersion(r.currentVersion)
		r.uncommittedEvents.add(evt)
	}

	return nil
}
// ListCommittedEvents returns a list of already committed events.
func (r *AggregateCluster) ListCommittedEvents() []event.Eventer {
	return r.committedEvents
}

// ListUncommittedEvents returns a list of not committed yet events.
func (r *AggregateCluster) ListUncommittedEvents() []event.Eventer {
	uncommittedEvents := make([]event.Eventer, 0, r.uncommittedEvents.len)
	r.uncommittedEvents.traverse(func(uncommitted event.Eventer) error {
		uncommittedEvents = append(uncommittedEvents, uncommitted)
		return nil
	})
	return uncommittedEvents
}

// Commit commits event and deletes from committed events list.
func (r *AggregateCluster) Commit(evt event.Eventer) error {
	r.uncommittedEvents.remove(evt)
	return nil
}

func (r *AggregateCluster) nextVersion() event.Version {
	return r.currentVersion + event.NextVersion
}

var ErrEventDuplication = errors.New("event duplication, event is already exist")

func (r *AggregateCluster) checkVersionDuplication(evt event.Eventer) error {
	for i := range r.committedEvents {
		committed := r.committedEvents[i]
		if committed.GetAggregateId() == evt.GetAggregateId() &&
			committed.GetAggregateType() == evt.GetAggregateType() &&
			committed.GetVersion() == evt.GetVersion() {

			return ErrEventDuplication
		}
	}
	return nil
}

func New(agg event.Aggregator, transition event.Transition, idgenfn IDGenerator) *AggregateCluster {
	return &AggregateCluster{
		currentId:         idgenfn(idDefaultAlphabet, idDefaultSize),
		currentType:       reflect.TypeOf(agg).Elem().Name(),
		committedEvents:   make([]event.Eventer, 0, 8),
		uncommittedEvents: new(linkedList),
		transitionfn:      transition,
	}
}