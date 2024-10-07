package test

import (
	top_properties "accommodation/genproto/top_properties"
	"accommodation/internal/config"
	logger "accommodation/internal/logs"
	"accommodation/internal/storage/postgres"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
)

func TestCreateTopProperty(t *testing.T) {
	cfg := config.Config{
		DB_PASSWORD: "salom",
		DB_HOST:     "localhost",
		DB_PORT:     5432,
		DB_USER:     "postgres",
		DB_NAME:     "rent_hub_accomodation",
	}

	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}
	defer db.Close()

	repo := postgres.NewTopPropertiesRepository(db, nil)

	req := &top_properties.CreateTopPropertyReq{
		PropertyId: "61d099f5-cad5-4e26-ac6c-1e85af32a7d9",
		UserId:     uuid.NewString(),
		TariffName: "Premium",
	}

	res, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Id == "" {
		t.Errorf("expected a valid top property ID, got empty string")
	}
}

func TestGetTopProperty(t *testing.T) {
	cfg := config.Config{
		DB_PASSWORD: "salom",
		DB_HOST:     "localhost",
		DB_PORT:     5432,
		DB_USER:     "postgres",
		DB_NAME:     "rent_hub_accomodation",
	}

	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}
	defer db.Close()

	repo := postgres.NewTopPropertiesRepository(db, nil)

	req := &top_properties.CreateTopPropertyReq{
		PropertyId: "7d0701fd-4f5a-4fa6-b214-42be11cca708",
		UserId:     uuid.NewString(),
		TariffName: "Standard",
	}

	createRes, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	getReq := &top_properties.GetTopPropertyReq{Id: createRes.Id}
	getRes, err := repo.Get(context.Background(), getReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if getRes.TopProperty.Id != createRes.Id {
		t.Errorf("expected top property ID %v, got %v", createRes.Id, getRes.TopProperty.Id)
	}
}

func TestDeleteTopProperty(t *testing.T) {
	cfg := config.Config{
		DB_PASSWORD: "salom",
		DB_HOST:     "localhost",
		DB_PORT:     5432,
		DB_USER:     "postgres",
		DB_NAME:     "rent_hub_accomodation",
	}

	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}
	defer db.Close()

	logs := logger.NewLogger()
	repo := postgres.NewTopPropertiesRepository(db, logs)

	req := &top_properties.CreateTopPropertyReq{
		PropertyId: "d37d9b42-90f9-441d-af74-f302556081c2",
		UserId:     uuid.NewString(),
		TariffName: "Deluxe",
	}

	createRes, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	deleteReq := &top_properties.DeleteTopPropertyReq{Id: createRes.Id}
	deleteRes, err := repo.Delete(context.Background(), deleteReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !deleteRes.Success {
		t.Errorf("expected success, got failure")
	}

	getReq := &top_properties.GetTopPropertyReq{Id: createRes.Id}
	_, err = repo.Get(context.Background(), getReq)
	if err == nil {
		t.Errorf("expected an error when fetching deleted top property, but got none")
	}
}

func TestGetAllTopProperties(t *testing.T) {
	cfg := config.Config{
		DB_PASSWORD: "salom",
		DB_HOST:     "localhost",
		DB_PORT:     5432,
		DB_USER:     "postgres",
		DB_NAME:     "rent_hub_accomodation",
	}

	db, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}
	defer db.Close()

	logs := logger.NewLogger()
	repo := postgres.NewTopPropertiesRepository(db, logs)

	for i := 0; i < 5; i++ {
		req := &top_properties.CreateTopPropertyReq{
			PropertyId: "d37d9b42-90f9-441d-af74-f302556081c2",
			UserId:     uuid.NewString(),
			TariffName: "Standard",
		}
		_, err := repo.Create(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected error creating top property: %v", err)
		}
	}

	getAllReq := &top_properties.GetAllTopPropertyReq{
		Page:  1,
		Limit: 3,
	}

	getAllRes, err := repo.GetAll(context.Background(), getAllReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(getAllRes.TopProperties) != 3 {
		t.Errorf("expected 3 top properties, got %v", len(getAllRes.TopProperties))
	}
}