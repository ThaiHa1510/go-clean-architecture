package valueobjects

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

type SerializerType string

const(
	SerializerTypeJson SerializerType = "json"
	SerializerTypeBSON SerializerType = "bson"
)

type Serializer interface {
	Encode(v interface{}) (EventPayload, error)
	Decode(data EventPayload, dst interface{}) error
}

var MatchedSerializers = map[SerializerType]Serializer{
	SerializerTypeJson: &JSONSerializer{},
	SerializerTypeBSON: &BSONSerializer{},
}

type JSONSerializer struct {
	
}

func (serial *JSONSerializer) Encode(v interface{}) EventPayload {
	return json.Marschal(v)
}

func (serial *JSONSerializer) Decode(data interface{}, dst interface{}) error {
	return json.Unmarshal(data, &dist)
}

type BSONSerializer struct{}

func (BSONSerializer) Encode(v interface{}) (Payload, error) {
	return bson.Marshal(v)
}

func (BSONSerializer) Decode(data Payload, dst interface{}) error {
	return bson.Unmarshal(data, &dst)
}

type UnsupportedSerializer struct{}

func (UnsupportedSerializer) Encode(v interface{}) (Payload, error) {
	return nil, errors.New("unsupported serializer")
}

func (UnsupportedSerializer) Decode(data Payload, dst interface{}) error {
	return errors.New("unsupported serializer")
}