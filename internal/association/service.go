package association

import (
	"fleet_management/proto"
)

type Service interface {
	AssociateDriverToVehicle(driverID int, vehicleID int) error
	GetAllAssociations() ([]*proto.Association, error)
	GetDriversByVehicle(vehicleID int) ([]int, error)
	GetVehiclesByDriver(driverID int) ([]int, error)
	DissociateDriverFromVehicle(driverID int, vehicleID int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) AssociateDriverToVehicle(driverID int, vehicleID int) error {
	return s.repo.AssociateDriverToVehicle(driverID, vehicleID)
}

func (s *service) GetAllAssociations() ([]*proto.Association, error) {
	return s.repo.GetAllAssociations()
}

func (s *service) GetDriversByVehicle(vehicleID int) ([]int, error) {
	return s.repo.GetDriversByVehicle(vehicleID)
}

func (s *service) GetVehiclesByDriver(driverID int) ([]int, error) {
	return s.repo.GetVehiclesByDriver(driverID)
}

func (s *service) DissociateDriverFromVehicle(driverID int, vehicleID int) error {
	return s.repo.DissociateDriverFromVehicle(driverID, vehicleID)
}
