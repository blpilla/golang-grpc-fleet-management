package vehicle

import (
	"fleet_management/proto"
)

type Service interface {
	Create(vehicle *proto.CreateVehicleRequest) (*proto.VehicleResponse, error)
	GetByID(id int32) (*proto.VehicleResponse, error)
	GetAll() ([]*proto.VehicleResponse, error)
	Update(vehicle *proto.UpdateVehicleRequest) (*proto.VehicleResponse, error)
	Delete(id int32) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(vehicle *proto.CreateVehicleRequest) (*proto.VehicleResponse, error) {
	return s.repo.Create(vehicle)
}

func (s *service) GetByID(id int32) (*proto.VehicleResponse, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAll() ([]*proto.VehicleResponse, error) {
	return s.repo.GetAll()
}

func (s *service) Update(vehicle *proto.UpdateVehicleRequest) (*proto.VehicleResponse, error) {
	return s.repo.Update(vehicle)
}

func (s *service) Delete(id int32) error {
	return s.repo.Delete(id)
}
