package main

import (
	"movies/handler"
	pb "movies/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("movies"),
		service.Version("3"),
	)

	// Register handler
	pb.RegisterMoviesHandler(srv.Server(), new(handler.Movies))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
