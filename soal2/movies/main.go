package main

import (
	"movies/handler"
	pb "movies/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	srv := service.New(
		service.Name("movies"),
		service.Version("3"),
	)

	pb.RegisterMoviesHandler(srv.Server(), new(handler.Movies))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
