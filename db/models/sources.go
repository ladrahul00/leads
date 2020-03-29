package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Sources typ definition
type Sources struct {
	SourceName string             `bson:"sourceName,omitempty"`
	SourceTag  string             `bson:"sourceTag,omitempty"`
	ID         primitive.ObjectID `bson:"_id,omitempty"`
}
