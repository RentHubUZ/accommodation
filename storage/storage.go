package storage

import (
	pb "accommodation/genproto/accommodation"
	"context"
)

type IStorage interface {
	User() IUserStorage
	Close()
}

type IUserStorage interface {
	CreateHouse(context.Context, *pb.CreateHouseReq) (*pb.CreateHouseRes, error)
	UpdateHouse(context.Context, *pb.UpdateHouseReq) (*pb.UpdateHouseRes, error)
	GetAllHouse(context.Context, *pb.GetallHouseReq) (*pb.GetAllHouseRes, error)
	GetByIdHouse(context.Context, *pb.GetByIdHouseReq) (*pb.GetByIdHouseRes, error)
	DeleteHouse(context.Context, *pb.DeleteHouseReq) (*pb.DeleteHouseRes, error)
}