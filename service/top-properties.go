package service

import (
	pb "accommodation/genproto/top-properties"
	"accommodation/storage"
	"accommodation/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type TopPropertiesService struct {
	pb.UnimplementedTopPropertiesServiceServer
	User storage.IStorage
	Log  *slog.Logger
}

func NewTopPropertiesService(db *sql.DB, log *slog.Logger) *TopPropertiesService {
	return &TopPropertiesService{
		User: postgres.NewPostgresStorage(db, log),
		Log:  log,
	}
}

func (s *TopPropertiesService) Get(ctx context.Context, req *pb.GetTopPropertyReq) (*pb.GetTopPropertyRes, error) {
	res, err := s.User.TopProperties().Get(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting top properties: %v", err.Error()))
		return nil, err
	}
	return res, err
}

func (s *TopPropertiesService) GetAll(ctx context.Context, req *pb.GetAllTopPropertyReq) (*pb.GetAllTopPropertyRes, error) {
	res, err := s.User.TopProperties().GetAll(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting all top properties: %v", err.Error()))
		return nil, err
	}
	return res, err
}

func (s *TopPropertiesService) Create(ctx context.Context, req *pb.CreateTopPropertyReq) (*pb.CreateTopPropertyRes, error) {
	res, err := s.User.TopProperties().Create(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error creating top property: %v", err.Error()))
		return nil, err
	}
	return res, err
}

func (s *TopPropertiesService) Delete(ctx context.Context, req *pb.DeleteTopPropertyReq) (*pb.DeleteTopPropertyRes, error) {
	res, err := s.User.TopProperties().Delete(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error deleting top property: %v", err.Error()))
		return nil, err
	}
	return res, err
}
