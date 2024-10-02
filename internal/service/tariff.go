package service

import (
	pb "accommodation/genproto/tariff"
	"accommodation/internal/storage"
	"accommodation/internal/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type TariffService struct {
	pb.UnimplementedTariffServiceServer
	User storage.IStorage
	Log  *slog.Logger
}

func NewTariffService(db *sql.DB, log *slog.Logger) *TariffService {
	return &TariffService{
		User: postgres.NewPostgresStorage(db, log),
		Log:  log,
	}
}

func (s *TariffService) Get(ctx context.Context, req *pb.GetTariffReq) (*pb.GetTariffRes, error) {
	res, err := s.User.Tariff().Get(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting tariff: %v", err.Error()))
		return nil, err
	}
	return res, err
}

func (s *TariffService) GetAll(ctx context.Context, req *pb.GetAllTariffReq) (*pb.GetAllTariffRes, error) {
	res, err := s.User.Tariff().GetAll(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting all tariffs: %v", err.Error()))
		return nil, err
	}
	return res, err
}

func (s *TariffService) Create(ctx context.Context, req *pb.CreateTariffReq) (*pb.CreateTariffRes, error) {
	res, err := s.User.Tariff().Create(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error creating tariff: %v", err.Error()))
		return nil, err
	}
	return res, err
}

func (s *TariffService) Update(ctx context.Context, req *pb.UpdateTariffReq) (*pb.UpdateTariffRes, error) {
	res, err := s.User.Tariff().Update(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error updating tariff: %v", err.Error()))
		return nil, err
	}
	return res, err
}

func (s *TariffService) Delete(ctx context.Context, req *pb.DeleteTariffReq) (*pb.DeleteTariffRes, error) {
	res, err := s.User.Tariff().Delete(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error deleting tariff: %v", err.Error()))
		return nil, err
	}
	return res, err
}
