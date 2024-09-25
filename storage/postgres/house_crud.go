package postgres

import (
	pb "accommodation/genproto/accommodation"
	"accommodation/pkg/logger"
	"accommodation/storage"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type HousesRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewHousesRepository(db *sql.DB) storage.IUserStorage {
	return &HousesRepository{Db: db, Log: logger.NewLogger()}
}

func (h *HousesRepository) CreateHouse(ctx context.Context, req *pb.CreateHouseReq) (*pb.CreateHouseRes, error) {
	query := `insert into houses (id, owner_id, title, description,
								 latitude, longitude, images_url, total_square_area,
								 price_per_month, bedrooms, bathrooms, created_at, updated_at
							)values(
								 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	id := uuid.NewString()
	newtime := time.Now()

	_,err := h.Db.Exec(query,id,req.OwnerId,req.Title,
						req.Description,req.Latitude,req.Longitude,
						req.ImagesUrl,req.TotalSquareArea,req.PricePerMonth,
						req.Bedrooms,req.Bathrooms,newtime,newtime)

	if err != nil {
		h.Log.Error(fmt.Sprintf("Error creating home: %v",err.Error()))
		return nil,err
	}

	return &pb.CreateHouseRes{
		Id: id,
		OwnerId: req.OwnerId,
		Title: req.Title,
		Description: req.Description,
		Latitude: req.Latitude,
		Longitude: req.Longitude,
		TotalSquareArea: req.TotalSquareArea,
		PricePerMonth: req.PricePerMonth,
		Bedrooms: req.Bedrooms,
		Bathrooms: req.Bathrooms,
		ImagesUrl: req.ImagesUrl,
		
	},nil
}

func (h *HousesRepository) UpdateHouse(ctx context.Context, req *pb.UpdateHouseReq) (*pb.UpdateHouseRes, error) {
	return nil, nil
}

func (h *HousesRepository) GetAllHouse(ctx context.Context, req *pb.GetallHouseReq) (*pb.GetAllHouseRes, error) {
	return nil, nil
}

func (h *HousesRepository) GetByIdHouse(ctx context.Context, req *pb.GetByIdHouseReq) (*pb.GetByIdHouseRes, error) {
	return nil, nil
}

func (h *HousesRepository) DeleteHouse(ctx context.Context, req *pb.DeleteHouseReq) (*pb.DeleteHouseRes, error) {
	return nil, nil
}
