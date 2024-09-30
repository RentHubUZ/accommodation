package service

import (
	pb "accommodation/genproto/payment"
	"accommodation/storage"
	"accommodation/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
	User storage.IStorage
	Log  *slog.Logger
}

func NewPaymentService(db *sql.DB, log *slog.Logger) *PaymentService {
	return &PaymentService{
		User: postgres.NewPostgresStorage(db, log),
		Log:  log,
	}
}

func (s *PaymentService) Create(ctx context.Context, req *pb.CreatePaymentReq) (*pb.CreatePaymentRes, error) {
	res, err := s.User.Payment().Create(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error creating payment: %v", err.Error()))
		return nil, err
	}
	return res, err
}

func (s *PaymentService) Get(ctx context.Context, req *pb.GetPaymentReq) (*pb.GetPaymentRes, error) {
	res, err := s.User.Payment().Get(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting payment: %v", err.Error()))
		return nil, err
	}	
	return res, err
}

func (s *PaymentService) GetAll(ctx context.Context, req *pb.GetAllPaymentReq) (*pb.GetAllPaymentRes, error) {
	res, err := s.User.Payment().GetAll(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting all payments: %v", err.Error()))
		return nil, err
	}
	return res, err
}