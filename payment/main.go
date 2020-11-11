package main

import (
	"payment/handler"
	pb "payment/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("payment"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPaymentHandler(srv.Server(), new(handler.Payment))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
