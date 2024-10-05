package test

import (
	pb "accommodation/genproto/accommodation"
	"accommodation/internal/config"
	"accommodation/internal/storage/postgres"
	"context"
	"fmt"
	"log"
	"testing"
)

func TestCreateHouse(t *testing.T) {
	cfg := config.Config{
		ACCOMMODATION_SERVICE: ":9999",
		DB_PASSWORD: "salom",
		DB_HOST: "localhost",
		DB_PORT: 5432,
		DB_USER: "postgres",
		DB_NAME: "rent_hub_accomodation",
	}
	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}
	repo := postgres.NewHousesRepository(db)

	req := &pb.CreateHouseReq{
		OwnerId:       "cef78672-3237-4d47-b0ae-0ea6b141e5c9",
		Address:       "123 Main St",
		Price:         1500,
		PropertyType:  "Apartment",
		Bedrooms:      3,
		Bathrooms:     2,
		SquareFootage: 1200,
		ListingStatus: "Available",
		Description:   "Beautiful house",
		RoommateCount: 0,
		LeaseTerms:    "12 months",
		LeaseDuration: 12,
		TopStatus:     false,
		Latitude:      41.1234,
		Longitude:     -75.1234,
		ImageUrl:      []string{"image1.jpg", "image2.jpg"},
	}

	res, err := repo.CreateHouse(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if res.OwnerId != req.OwnerId || res.Address != req.Address {
		t.Errorf("unexpected result: got %v", res)
	}
}

func TestUpdateHouse(t *testing.T) {
	cfg := config.Config{
		ACCOMMODATION_SERVICE: ":9999",
		DB_PASSWORD: "salom",
		DB_HOST: "localhost",
		DB_PORT: 5432,
		DB_USER: "postgres",
		DB_NAME: "rent_hub_accomodation",
	}
	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}
	repo := postgres.NewHousesRepository(db)

	req := &pb.UpdateHouseReq{
		Id:            "d37d9b42-90f9-441d-af74-f302556081c2",
		Price:         1600,
		PropertyType:  "House",
		Bedrooms:      4,
		Bathrooms:     3,
		SquareFootage: 1400,
		Description:   "Updated house",
	}

	res, err := repo.UpdateHouse(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	fmt.Printf("resp: %v\n", res)
}

func TestGetHouseById(t *testing.T) {
	cfg := config.Config{
		ACCOMMODATION_SERVICE: ":9999",
		DB_PASSWORD: "salom",
		DB_HOST: "localhost",
		DB_PORT: 5432,
		DB_USER: "postgres",
		DB_NAME: "rent_hub_accomodation",
	}
	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}
	repo := postgres.NewHousesRepository(db)

	req := &pb.GetByIdHouseReq{
		Id: "d37d9b42-90f9-441d-af74-f302556081c2",
	}

	res, err := repo.GetByIdHouse(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if res.Id != req.Id {
		t.Errorf("expected ID %v, got %v", req.Id, res.Id)
	}
}

func TestGetAllHouses(t *testing.T) {
	cfg := config.Config{
		ACCOMMODATION_SERVICE: ":9999",
		DB_PASSWORD: "salom",
		DB_HOST: "localhost",
		DB_PORT: 5432,
		DB_USER: "postgres",
		DB_NAME: "rent_hub_accomodation",
	}
	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}
	repo := postgres.NewHousesRepository(db)

	req := &pb.GetallHouseReq{
		Limit: 10,
		Page:  0,
	}

	res, err := repo.GetAllHouse(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(res.House) == 0 {
		t.Errorf("expected some houses, but got none")
	}

	for _, house := range res.House {
		if house.Address == "" || house.OwnerId == "" {
			t.Errorf("unexpected empty fields in house: %v", house)
		}
	}
}

func TestDeleteHouse(t *testing.T) {
	cfg := config.Config{
		ACCOMMODATION_SERVICE: ":9999",
		DB_PASSWORD: "salom",
		DB_HOST: "localhost",
		DB_PORT: 5432,
		DB_USER: "postgres",
		DB_NAME: "rent_hub_accomodation",
	}
	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}

	repo := postgres.NewHousesRepository(db)

	req := &pb.DeleteHouseReq{
		Id: "d37d9b42-90f9-441d-af74-f302556081c2",
	}

	res, err := repo.DeleteHouse(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	fmt.Println(res)
}