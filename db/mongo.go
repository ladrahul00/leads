package db

import (
	"context"

	"leads/constants"

	log "github.com/micro/go-micro/v2/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(constants.MongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to MongoDB!...")

	return client.Database(constants.DatabaseName)
}

// Leads is used to retrieve collection from database
func Leads() *mongo.Collection {
	return connectMongo().Collection(constants.Leads)
}

// Campaigns used to retrieve collection from database
func Campaigns() *mongo.Collection {
	return connectMongo().Collection(constants.Campaigns)
}

// LeadTemplates used to retrieve collection from database
func LeadTemplates() *mongo.Collection {
	return connectMongo().Collection(constants.LeadTemplates)
}

// Sources used to retrieve collection from database
func Sources() *mongo.Collection {
	return connectMongo().Collection(constants.Sources)
}
