package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Campaigns typ definition
type Campaigns struct {
	CampaignName string             `bson:"campaignName,omitempty"`
	CampaignTag  string             `bson:"campaignTag,omitempty"`
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	TemplateID   primitive.ObjectID `bson:"templateId,omitempty"`
}
