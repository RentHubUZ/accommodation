package service

import (
	pb "accommodation/genproto/accommodation"
	"accommodation/storage"
	"accommodation/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type HouseService struct {
	pb.UnimplementedAccommodationServiceServer
	User storage.IStorage
	Log  *slog.Logger
}

func NewHouseService(db *sql.DB, log *slog.Logger) *HouseService {
	return &HouseService{
		User: postgres.NewPostgresStorage(db, log),
		Log:  log,
	}
}

func (s *HouseService) CreateHouse(ctx context.Context, req *pb.CreateHouseReq) (*pb.CreateHouseRes, error) {
	res, err := s.User.House().CreateHouse(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error creating house service: %v",err.Error()))
		return nil,err
	}

	return res,nil
}

func (s *HouseService) UpdateHouse(ctx context.Context,req *pb.UpdateHouseReq) (*pb.UpdateHouseRes, error) {
	res,err := s.User.House().UpdateHouse(ctx,req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("error updating house service: %v",err.Error()))
		return nil,err
	}
	return res,nil
}

func (s *HouseService) GetAllHouse(ctx context.Context,req *pb.GetallHouseReq) (*pb.GetAllHouseRes, error) {
	res,err := s.User.House().GetAllHouse(ctx,req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting all data from house service: %v",err.Error()))
		return nil,err
	}

	return res,nil
}

func (s *HouseService) GetByIdHouse(ctx context.Context,req *pb.GetByIdHouseReq) (*pb.GetByIdHouseRes, error) {
	res,err:=s.User.House().GetByIdHouse(ctx,req)
	if err!=nil{
		s.Log.Error(fmt.Sprintf("Error retrieving id information from house service: %v",err.Error()))
		return nil,err
	}
	return res,nil
}

func (s *HouseService) DeleteHouse(ctx context.Context,req *pb.DeleteHouseReq) (*pb.DeleteHouseRes, error){
	res,err:=s.User.House().DeleteHouse(ctx,req)
	if err != nil{
		s.Log.Error(fmt.Sprintf("Error deleting reference in house service: %v",err))
		return nil,err
	}
	return res,nil
}