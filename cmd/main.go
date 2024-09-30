package main

import (
	"accommodation/internal/config"
	"accommodation/internal/logs"
	"accommodation/service"
	"accommodation/storage/postgres"
	"log"
	"net"
	ph "accommodation/genproto/accommodation"
	pt "accommodation/genproto/tariff"
	top "accommodation/genproto/top-properties"
	pay "accommodation/genproto/payment"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	config := config.Load()
	listener, err := net.Listen("tcp", config.ACCOMMODATION_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	logs := logger.NewLogger()

	db, err := postgres.ConnectionDb(&config)
	if err != nil {
		log.Fatal(err)
	}

	HouseService := service.NewHouseService(db, logs)
	PaymentService := service.NewPaymentService(db, logs)
	TopService := service.NewTopPropertiesService(db, logs)
	TariffService := service.NewTariffService(db, logs)


	server := grpc.NewServer()
	ph.RegisterAccommodationServiceServer(server, HouseService)
	pt.RegisterTariffServiceServer(server, TariffService)
	top.RegisterTopPropertiesServiceServer(server, TopService)
	pay.RegisterPaymentServiceServer(server, PaymentService)

	log.Printf("Server is listening on port %s\n", config.ACCOMMODATION_SERVICE)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}