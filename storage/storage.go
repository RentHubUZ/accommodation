package storage

import (
	pb "accommodation/genproto/accommodation"
	"context"
)

type IStorage interface {
	User() IHouseStorage
	Close()
}

type IHouseStorage interface {
	CreateHouse(ctx context.Context,req *pb.CreateHouseReq) (*pb.CreateHouseRes, error)
	UpdateHouse(ctx context.Context,req *pb.UpdateHouseReq) (*pb.UpdateHouseRes, error)
	GetAllHouse(ctx context.Context,req *pb.GetallHouseReq) ([]*pb.GetAllHouseRes, error)
	GetByIdHouse(ctx context.Context,req *pb.GetByIdHouseReq) (*pb.GetByIdHouseRes, error)
	DeleteHouse(ctx context.Context,req *pb.DeleteHouseReq) (*pb.DeleteHouseRes, error)
}