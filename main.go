package main

import (
	"leads/handler"
	leads "leads/proto/leads"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.leads"),
		micro.Version("0.1"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	leads.RegisterLeadsHandler(service.Server(), new(handler.LeadsRequestHandler))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.service.leads", service.Server(), new(subscriber.Leads))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
