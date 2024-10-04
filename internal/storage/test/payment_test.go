package test

import (
	"accommodation/genproto/payment"
	"accommodation/internal/config"
	"accommodation/internal/storage/postgres"
	"context"
	"log"
	"testing"
)

func TestCreatePayment(t *testing.T) {
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

	repo := postgres.NewPaymentRepository(db, nil)

	req := &payment.CreatePaymentReq{
		Amount: "1500.50",
		Status: "Completed",
	}

	res, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Id == "" {
		t.Errorf("expected a valid payment ID, got empty string")
	}
}

func TestGetPayment(t *testing.T) {
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

	repo := postgres.NewPaymentRepository(db, nil)

	req := &payment.CreatePaymentReq{
		Amount: "1200.00",
		Status: "Pending",
	}

	createRes, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	getReq := &payment.GetPaymentReq{Id: createRes.Id}
	getRes, err := repo.Get(context.Background(), getReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if getRes.Payment.Id != createRes.Id {
		t.Errorf("expected payment ID %v, got %v", createRes.Id, getRes.Payment.Id)
	}
}

func TestGetAllPayments(t *testing.T) {
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

	repo := postgres.NewPaymentRepository(db, nil)

	for i := 0; i < 5; i++ {
		req := &payment.CreatePaymentReq{
			Amount: "1000.00",
			Status: "Completed",
		}
		_, err := repo.Create(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected error creating payment: %v", err)
		}
	}

	getAllReq := &payment.GetAllPaymentReq{
		Page:  1,
		Limit: 3, 
	}

	getAllRes, err := repo.GetAll(context.Background(), getAllReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(getAllRes.Payments) != 3 {
		t.Errorf("expected 3 payments, got %v", len(getAllRes.Payments))
	}
}

func TestDeletePayment(t *testing.T) {
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

	repo := postgres.NewPaymentRepository(db, nil)

	req := &payment.CreatePaymentReq{
		Amount: "1500.00",
		Status: "Pending",
	}

	createRes, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	deleteReq := &payment.DeletePaymentReq{Id: createRes.Id}
	deleteRes, err := repo.Delete(context.Background(), deleteReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !deleteRes.Success {
		t.Errorf("expected success, got failure")
	}

	getReq := &payment.GetPaymentReq{Id: createRes.Id}
	_, err = repo.Get(context.Background(), getReq)
	if err == nil {
		t.Errorf("expected an error when fetching deleted payment, but got none")
	}
}
