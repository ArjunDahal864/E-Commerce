package main

import (
	"recommendations/handler"
	pb "recommendations/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("recommendations"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterRecommendationsHandler(srv.Server(), new(handler.Recommendations))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
