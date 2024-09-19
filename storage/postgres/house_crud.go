package postgres

import (
	"accommodation/pkg/logger"
	"accommodation/storage"
	"context"
	"database/sql"
	"log/slog"
	pb "accommodation/genproto/accommodation"
)

type HousesRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewHousesRepository(db *sql.DB) storage.IUserStorage {
	return &HousesRepository{Db: db, Log: logger.NewLogger()}
}

func (h *HousesRepository) CreateHouse(context.Context, *pb.CreateHouseReq) (*pb.CreateHouseRes, error){
	return nil,nil
}

func (h *HousesRepository) UpdateHouse(context.Context, *pb.UpdateHouseReq) (*pb.UpdateHouseRes, error){
	return nil,nil
}

func (h *HousesRepository) GetAllHouse(context.Context, *pb.GetallHouseReq) (*pb.GetAllHouseRes, error){
	return nil,nil
}

func (h *HousesRepository) GetByIdHouse(context.Context, *pb.GetByIdHouseReq) (*pb.GetByIdHouseRes, error){
	return nil,nil
}

func (h *HousesRepository) DeleteHouse(context.Context, *pb.DeleteHouseReq) (*pb.DeleteHouseRes, error){
	return nil,nil
}