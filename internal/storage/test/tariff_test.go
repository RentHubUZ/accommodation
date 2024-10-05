package test

import (
	"accommodation/genproto/tariff"
	"accommodation/internal/config"
	logger "accommodation/internal/logs"
	"accommodation/internal/storage/postgres"
	"context"
	"log"
	"strconv"
	"testing"
)

func TestCreateTariff(t *testing.T) {
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
	repo := postgres.NewTariffsRepository(db, logs)

	for i := 0; i < 5; i++ {
		sonstr := strconv.Itoa(i)
		req := &tariff.CreateTariffReq{
			Name:   "Test Tariff " + sonstr, 
			Days:   30,
			Price:  float32(i * 100),
			Offers: "Some Offer",
		}
		_, err := repo.Create(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected error creating tariff: %v", err)
		}
	}
}

func TestGetTariff(t *testing.T) {
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

	repo := postgres.NewTariffsRepository(db, nil)

	req := &tariff.CreateTariffReq{
		Name:   "Standard",
		Days:   15,
		Price:  50.00,
		Offers: "Limited Listings",
	}

	createRes, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	getReq := &tariff.GetTariffReq{Id: createRes.Id}
	getRes, err := repo.Get(context.Background(), getReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if getRes.Tariff.Id != createRes.Id {
		t.Errorf("expected tariff ID %v, got %v", createRes.Id, getRes.Tariff.Id)
	}
}

func TestUpdateTariff(t *testing.T) {
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

	repo := postgres.NewTariffsRepository(db, nil)

	req := &tariff.CreateTariffReq{
		Name:   "Basic",
		Days:   7,
		Price:  20.00,
		Offers: "Basic Listings",
	}

	createRes, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updateReq := &tariff.UpdateTariffReq{
		Id:     createRes.Id,
		Name:   "Basic Plus",
		Days:   10,
		Price:  30.00,
		Offers: "More Listings",
	}

	updateRes, err := repo.Update(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !updateRes.Success {
		t.Errorf("expected update success, got failure")
	}

	getReq := &tariff.GetTariffReq{Id: createRes.Id}
	getRes, err := repo.Get(context.Background(), getReq)
	if err != nil {
		t.Fatalf("unexpected error retrieving updated tariff: %v", err)
	}

	if getRes.Tariff.Name != "Basic Plus" || getRes.Tariff.Days != 10 {
		t.Errorf("expected updated name and days, got %v and %v", getRes.Tariff.Name, getRes.Tariff.Days)
	}
}

func TestDeleteTariff(t *testing.T) {
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
	repo := postgres.NewTariffsRepository(db, logs)

	req := &tariff.CreateTariffReq{
		Name:   "Deluxe",
		Days:   60,
		Price:  200.00,
		Offers: "Exclusive Listings",
	}

	createRes, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	deleteReq := &tariff.DeleteTariffReq{Id: createRes.Id}
	deleteRes, err := repo.Delete(context.Background(), deleteReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !deleteRes.Success {
		t.Errorf("expected success, got failure")
	}

	getReq := &tariff.GetTariffReq{Id: createRes.Id}
	_, err = repo.Get(context.Background(), getReq)
	if err == nil {
		t.Errorf("expected an error when fetching deleted tariff, but got none")
	}
}

func TestGetAllTariffs(t *testing.T) {
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

	repo := postgres.NewTariffsRepository(db, nil)

	for i := 0; i < 5; i++ {
		req := &tariff.CreateTariffReq{
			Name:   "Test Tariff " + strconv.Itoa(i),
			Days:   30,
			Price:  float32(100 * i),
			Offers: "Test Offers",
		}
		_, err := repo.Create(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected error creating tariff: %v", err)
		}
	}

	getAllReq := &tariff.GetAllTariffReq{
		Page:  1,
		Limit: 3,
	}

	getAllRes, err := repo.GetAll(context.Background(), getAllReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(getAllRes.Tariffs) != 3 {
		t.Errorf("expected 3 tariffs, got %v", len(getAllRes.Tariffs))
	}
}
