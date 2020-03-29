package services

import (
	"context"
	"fmt"
	"leads/db"
	"leads/db/models"

	"github.com/micro/go-micro/util/log"
	"go.mongodb.org/mongo-driver/bson"
)

func sourceByTag(ctx context.Context, tag string) (models.Sources, error) {
	helper := db.Sources()
	filter := bson.M{"sourceTag": tag}

	var source models.Sources

	err := helper.FindOne(ctx, filter).Decode(&source)

	if err != nil {
		log.Error(err)
		return source, fmt.Errorf("source with the %s tag is not available", tag)
	}

	return source, nil
}
