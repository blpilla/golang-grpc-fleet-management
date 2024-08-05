package vehicle

import (
	"errors"
	"fleet_management/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	vehicles map[int32]*proto.VehicleResponse
}

func newMockRepository() *mockRepository {
	return &mockRepository{vehicles: make(map[int32]*proto.VehicleResponse)}
}

func (m *mockRepository) Create(vehicle *proto.CreateVehicleRequest) (*proto.VehicleResponse, error) {
	id := int32(len(m.vehicles) + 1)
	newVehicle := &proto.VehicleResponse{
		Id:    id,
		Make:  vehicle.Make,
		Model: vehicle.Model,
	}
	m.vehicles[id] = newVehicle
	return newVehicle, nil
}

func (m *mockRepository) GetByID(id int32) (*proto.VehicleResponse, error) {
	if vehicle, exists := m.vehicles[id]; exists {
		return vehicle, nil
	}
	return nil, errors.New("vehicle not found")
}

func (m *mockRepository) GetAll() ([]*proto.VehicleResponse, error) {
	var vehicles []*proto.VehicleResponse
	for _, vehicle := range m.vehicles {
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func (m *mockRepository) Update(vehicle *proto.UpdateVehicleRequest) (*proto.VehicleResponse, error) {
	if _, exists := m.vehicles[vehicle.Id]; !exists {
		return nil, errors.New("vehicle not found")
	}
	updatedVehicle := &proto.VehicleResponse{
		Id:    vehicle.Id,
		Make:  vehicle.Make,
		Model: vehicle.Model,
	}
	m.vehicles[vehicle.Id] = updatedVehicle
	return updatedVehicle, nil
}

func (m *mockRepository) Delete(id int32) error {
	if _, exists := m.vehicles[id]; !exists {
		return errors.New("vehicle not found")
	}
	delete(m.vehicles, id)
	return nil
}

func TestVehicleService(t *testing.T) {
	mockRepo := newMockRepository()
	service := NewService(mockRepo)

	t.Run("CreateVehicle", func(t *testing.T) {
		vehicle := &proto.CreateVehicleRequest{
			Make:  "Toyota",
			Model: "Camry",
		}
		createdVehicle, err := service.Create(vehicle)
		assert.NoError(t, err)
		assert.Equal(t, "Toyota", createdVehicle.Make)
		assert.Equal(t, "Camry", createdVehicle.Model)
	})

	t.Run("GetVehicle", func(t *testing.T) {
		vehicle, err := service.GetByID(1)
		assert.NoError(t, err)
		assert.Equal(t, "Toyota", vehicle.Make)
		assert.Equal(t, "Camry", vehicle.Model)
	})

	t.Run("UpdateVehicle", func(t *testing.T) {
		vehicle := &proto.UpdateVehicleRequest{
			Id:    1,
			Make:  "Honda",
			Model: "Accord",
		}
		updatedVehicle, err := service.Update(vehicle)
		assert.NoError(t, err)
		assert.Equal(t, "Honda", updatedVehicle.Make)
		assert.Equal(t, "Accord", updatedVehicle.Model)
	})

	t.Run("DeleteVehicle", func(t *testing.T) {
		err := service.Delete(1)
		assert.NoError(t, err)

		_, err = service.GetByID(1)
		assert.Error(t, err)
	})
}
