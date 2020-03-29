package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LeadTemplate type definition
type LeadTemplate struct {
	ID            primitive.ObjectID      `bson:"_id,omitempty"`
	Name          string                  `bson:"name"`
	CreatedBy     primitive.ObjectID      `bson:"createdBy"`
	KeyValueTypes []templateKeyValueTypes `bson:"keyValueTypes,omitempty"`
}

type templateKeyValueTypes struct {
	Key       string `bson:"key,omitempty"`
	valueType string `bson:"valueType,omitempty"`
}
