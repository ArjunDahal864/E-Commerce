package main

import (
	"notification/handler"
	pb "notification/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("notification"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterNotificationHandler(srv.Server(), new(handler.Notification))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
