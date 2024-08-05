package vehicle

import (
	"context"

	"fleet_management/proto"
)

type grpcService struct {
	service Service
	proto.UnimplementedVehicleServiceServer
}

func NewGRPCService(service Service) proto.VehicleServiceServer {
	return &grpcService{service: service}
}

func (s *grpcService) CreateVehicle(ctx context.Context, req *proto.CreateVehicleRequest) (*proto.VehicleResponse, error) {
	return s.service.Create(req)
}

func (s *grpcService) GetVehicleByID(ctx context.Context, req *proto.GetVehicleByIDRequest) (*proto.VehicleResponse, error) {
	return s.service.GetByID(req.GetId())
}

func (s *grpcService) GetAllVehicles(ctx context.Context, req *proto.GetAllVehiclesRequest) (*proto.GetAllVehiclesResponse, error) {
	vehicles, err := s.service.GetAll()
	if err != nil {
		return nil, err
	}
	return &proto.GetAllVehiclesResponse{Vehicles: vehicles}, nil
}

func (s *grpcService) UpdateVehicle(ctx context.Context, req *proto.UpdateVehicleRequest) (*proto.VehicleResponse, error) {
	return s.service.Update(req)
}

func (s *grpcService) DeleteVehicle(ctx context.Context, req *proto.DeleteVehicleRequest) (*proto.DeleteVehicleResponse, error) {
	err := s.service.Delete(req.GetId())
	return &proto.DeleteVehicleResponse{}, err
}
