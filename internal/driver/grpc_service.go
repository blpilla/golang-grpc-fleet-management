package driver

import (
	"context"

	"fleet_management/proto"
)

type grpcService struct {
	service Service
	proto.UnimplementedDriverServiceServer
}

func NewGRPCService(service Service) proto.DriverServiceServer {
	return &grpcService{service: service}
}

func (s *grpcService) CreateDriver(ctx context.Context, req *proto.CreateDriverRequest) (*proto.DriverResponse, error) {
	driver, err := s.service.Create(req)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (s *grpcService) GetDriverByID(ctx context.Context, req *proto.GetDriverByIDRequest) (*proto.DriverResponse, error) {
	driver, err := s.service.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (s *grpcService) GetAllDrivers(ctx context.Context, req *proto.GetAllDriversRequest) (*proto.GetAllDriversResponse, error) {
	drivers, err := s.service.GetAll()
	if err != nil {
		return nil, err
	}
	return &proto.GetAllDriversResponse{Drivers: drivers}, nil
}

func (s *grpcService) UpdateDriver(ctx context.Context, req *proto.UpdateDriverRequest) (*proto.DriverResponse, error) {
	driver, err := s.service.Update(req)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (s *grpcService) DeleteDriver(ctx context.Context, req *proto.DeleteDriverRequest) (*proto.DeleteDriverResponse, error) {
	err := s.service.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteDriverResponse{}, nil
}
