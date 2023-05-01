package main

import (
	"go-grpc-services-course/internal/db"
	"go-grpc-services-course/internal/rocket"
	"go-grpc-services-course/internal/transport/grpc"
	"log"
)

func Run() error {
	rocketStore, err := db.New()

	if err != nil {
		return err
	}

	err = rocketStore.Migrate()

	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	rktService := rocket.New(rocketStore)

	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
