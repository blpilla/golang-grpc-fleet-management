package driver

import (
	"errors"
	"fleet_management/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	drivers map[int32]*proto.DriverResponse
}

func newMockRepository() *mockRepository {
	return &mockRepository{drivers: make(map[int32]*proto.DriverResponse)}
}

func (m *mockRepository) Create(driver *proto.CreateDriverRequest) (*proto.DriverResponse, error) {
	id := int32(len(m.drivers) + 1)
	newDriver := &proto.DriverResponse{
		Id:        id,
		FirstName: driver.FirstName,
		LastName:  driver.LastName,
		License:   driver.License,
	}
	m.drivers[id] = newDriver
	return newDriver, nil
}

func (m *mockRepository) GetByID(id int32) (*proto.DriverResponse, error) {
	if driver, exists := m.drivers[id]; exists {
		return driver, nil
	}
	return nil, errors.New("driver not found")
}

func (m *mockRepository) GetAll() ([]*proto.DriverResponse, error) {
	var drivers []*proto.DriverResponse
	for _, driver := range m.drivers {
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (m *mockRepository) Update(driver *proto.UpdateDriverRequest) (*proto.DriverResponse, error) {
	if _, exists := m.drivers[driver.Id]; !exists {
		return nil, errors.New("driver not found")
	}
	updatedDriver := &proto.DriverResponse{
		Id:        driver.Id,
		FirstName: driver.FirstName,
		LastName:  driver.LastName,
		License:   driver.License,
	}
	m.drivers[driver.Id] = updatedDriver
	return updatedDriver, nil
}

func (m *mockRepository) Delete(id int32) error {
	if _, exists := m.drivers[id]; !exists {
		return errors.New("driver not found")
	}
	delete(m.drivers, id)
	return nil
}

func TestDriverService(t *testing.T) {
	mockRepo := newMockRepository()
	service := NewService(mockRepo)

	t.Run("CreateDriver", func(t *testing.T) {
		driver := &proto.CreateDriverRequest{
			FirstName: "John",
			LastName:  "Doe",
			License:   "ABC123",
		}
		createdDriver, err := service.Create(driver)
		assert.NoError(t, err)
		assert.Equal(t, "John", createdDriver.FirstName)
		assert.Equal(t, "Doe", createdDriver.LastName)
		assert.Equal(t, "ABC123", createdDriver.License)
	})

	t.Run("GetDriver", func(t *testing.T) {
		driver, err := service.GetByID(1)
		assert.NoError(t, err)
		assert.Equal(t, "John", driver.FirstName)
		assert.Equal(t, "Doe", driver.LastName)
	})

	t.Run("UpdateDriver", func(t *testing.T) {
		driver := &proto.UpdateDriverRequest{
			Id:        1,
			FirstName: "Jane",
			LastName:  "Doe",
			License:   "XYZ789",
		}
		updatedDriver, err := service.Update(driver)
		assert.NoError(t, err)
		assert.Equal(t, "Jane", updatedDriver.FirstName)
		assert.Equal(t, "Doe", updatedDriver.LastName)
		assert.Equal(t, "XYZ789", updatedDriver.License)
	})

	t.Run("DeleteDriver", func(t *testing.T) {
		err := service.Delete(1)
		assert.NoError(t, err)

		_, err = service.GetByID(1)
		assert.Error(t, err)
	})
}
