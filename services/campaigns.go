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

// CampaignIDByTag retrieves campaign id by campaign tag
func campaignByTag(ctx context.Context, tag string) (models.Campaigns, error) {
	var campaign models.Campaigns
	helper := db.Campaigns()
	filter := bson.M{"campaignTag": tag}
	err := helper.FindOne(ctx, filter).Decode(&campaign)
	if err != nil {
		log.Error(err)
		return campaign, fmt.Errorf("campaign with the %s tag is not available", tag)
	}
	return campaign, nil
}

// CampaignByID retrieve
func campaignByID(ctx context.Context, id primitive.ObjectID) (models.Campaigns, error) {
	var campaign models.Campaigns
	helper := db.Campaigns()
	filter := bson.M{"_id": id}
	err := helper.FindOne(ctx, filter).Decode(&campaign)
	if err != nil {
		log.Error(err)
		return campaign, fmt.Errorf("campaign with the %s id is not available", id.String())
	}
	return campaign, nil
}
