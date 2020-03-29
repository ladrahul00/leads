package services

import (
	"context"
	"fmt"
	"leads/db"
	"leads/db/models"

	"github.com/micro/go-micro/debug/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LeadTeamplateByCampaignID retrieve
func leadTeamplateByCampaignID(ctx context.Context, campaignID primitive.ObjectID) (models.LeadTemplate, error) {
	var leadTemplate models.LeadTemplate
	campaign, err := campaignByID(ctx, campaignID)
	if err != nil {
		return leadTemplate, err
	}

	helper := db.LeadTemplates()
	filter := bson.M{"_id": campaign.TemplateID}
	err = helper.FindOne(ctx, filter).Decode(&leadTemplate)
	if err != nil {
		log.Error(err)
		return leadTemplate, fmt.Errorf("lead template with the templateId %s is not available", campaign.TemplateID.String())
	}
	return leadTemplate, nil
}
