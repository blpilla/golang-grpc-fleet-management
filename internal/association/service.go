package association

import "fleet_management/proto"

type Service interface {
	AssociateDriverToVehicle(driverID uint32, vehicleID uint32) error
	GetAllAssociations() ([]*proto.Association, error)
	GetDriversByVehicle(vehicleID uint32) ([]uint32, error)
	GetVehiclesByDriver(driverID uint32) ([]uint32, error)
	DissociateDriverFromVehicle(driverID uint32, vehicleID uint32) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) AssociateDriverToVehicle(driverID uint32, vehicleID uint32) error {
	return s.repo.AssociateDriverToVehicle(driverID, vehicleID)
}

func (s *service) GetAllAssociations() ([]*proto.Association, error) {
	associations, err := s.repo.GetAllAssociations()
	if err != nil {
		return nil, err
	}

	var protoAssociations []*proto.Association
	for _, assoc := range associations {
		protoAssociations = append(protoAssociations, &proto.Association{
			DriverId:  int32(assoc.DriverID),
			VehicleId: int32(assoc.VehicleID),
		})
	}
	return protoAssociations, nil
}

func (s *service) GetDriversByVehicle(vehicleID uint32) ([]uint32, error) {
	return s.repo.GetDriversByVehicle(vehicleID)
}

func (s *service) GetVehiclesByDriver(driverID uint32) ([]uint32, error) {
	return s.repo.GetVehiclesByDriver(driverID)
}

func (s *service) DissociateDriverFromVehicle(driverID uint32, vehicleID uint32) error {
	return s.repo.DissociateDriverFromVehicle(driverID, vehicleID)
}
