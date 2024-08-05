package association

import (
	"fleet_management/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAssociationRepository struct {
	associations map[int32]*proto.Association
}

func newMockAssociationRepository() *mockAssociationRepository {
	return &mockAssociationRepository{associations: make(map[int32]*proto.Association)}
}

func (m *mockAssociationRepository) AssociateDriverToVehicle(driverID, vehicleID int) error {
	association := &proto.Association{DriverId: int32(driverID), VehicleId: int32(vehicleID)}
	m.associations[int32(driverID)] = association
	return nil
}

func (m *mockAssociationRepository) DissociateDriverFromVehicle(driverID, vehicleID int) error {
	delete(m.associations, int32(driverID))
	return nil
}

func (m *mockAssociationRepository) GetAllAssociations() ([]*proto.Association, error) {
	associations := []*proto.Association{}
	for _, assoc := range m.associations {
		associations = append(associations, assoc)
	}
	return associations, nil
}

func (m *mockAssociationRepository) GetDriversByVehicle(vehicleID int) ([]int, error) {
	var driverIDs []int
	for _, assoc := range m.associations {
		if int(assoc.VehicleId) == vehicleID {
			driverIDs = append(driverIDs, int(assoc.DriverId))
		}
	}
	return driverIDs, nil
}

func (m *mockAssociationRepository) GetVehiclesByDriver(driverID int) ([]int, error) {
	var vehicleIDs []int
	for _, assoc := range m.associations {
		if int(assoc.DriverId) == driverID {
			vehicleIDs = append(vehicleIDs, int(assoc.VehicleId))
		}
	}
	return vehicleIDs, nil
}

func TestAssociationService(t *testing.T) {
	mockRepo := newMockAssociationRepository()
	service := NewService(mockRepo)

	t.Run("AssociateDriverToVehicle", func(t *testing.T) {
		err := service.AssociateDriverToVehicle(1, 1)
		assert.NoError(t, err)
		association := mockRepo.associations[1]
		assert.Equal(t, int32(1), association.DriverId)
		assert.Equal(t, int32(1), association.VehicleId)
	})
}
