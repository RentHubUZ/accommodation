package main

import (
	"accommodation/config"
	"accommodation/pkg/logger"
	"accommodation/service"
	"accommodation/storage/postgres"
	"log"
	"net"
	pb "accommodation/genproto/accommodation"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().ACCOMMODATION_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal(err)
	}

	logs := logger.NewLogger()
	service := service.NewHouseService(db, logs)

	server := grpc.NewServer()
	pb.RegisterAccommodationServiceServer(server, service)

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}