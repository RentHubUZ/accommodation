package storage

import (
	pb "accommodation/genproto/accommodation"
	pt "accommodation/genproto/tariff"
	"context"
)

type IStorage interface {
	House() IHouseStorage
	Tariff() ITariffStorage
	Close()
}

type IHouseStorage interface {
	CreateHouse(ctx context.Context,req *pb.CreateHouseReq) (*pb.CreateHouseRes, error)
	UpdateHouse(ctx context.Context,req *pb.UpdateHouseReq) (*pb.UpdateHouseRes, error)
	GetAllHouse(ctx context.Context,req *pb.GetallHouseReq) (*pb.GetAllHouseRes, error)
	GetByIdHouse(ctx context.Context,req *pb.GetByIdHouseReq) (*pb.GetByIdHouseRes, error)
	DeleteHouse(ctx context.Context,req *pb.DeleteHouseReq) (*pb.DeleteHouseRes, error)
}

type ITariffStorage interface {
	Create(ctx context.Context, req *pt.CreateTariffReq) (*pt.CreateTariffRes, error)
	Get(ctx context.Context, req *pt.GetTariffReq) (*pt.GetTariffRes, error)
	GetAll(ctx context.Context, req *pt.GetAllTariffReq) (*pt.GetAllTariffRes, error)
	Update(ctx context.Context, req *pt.UpdateTariffReq) (*pt.UpdateTariffRes, error)
	Delete(ctx context.Context, req *pt.DeleteTariffReq) (*pt.DeleteTariffRes, error)
}
