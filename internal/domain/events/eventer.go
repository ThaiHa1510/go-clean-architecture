package event

import(
	"github.com/go-clean-archituecture/valueobjects"
)

type Eventer interface {
	GetAggregateId() string
	SetAggregateId(id string)
	GetAggregateType() string
	SetAggregateType(type string)
	GetReason() string
	SetReason(reason string)
	GetVersion() int
	SetVersion(version int)
	GetTimestamp() valueobjects.Timestamp
}