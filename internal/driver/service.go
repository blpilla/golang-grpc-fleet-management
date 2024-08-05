package driver

import (
	"fleet_management/proto"
)

type Service interface {
	Create(driver *proto.CreateDriverRequest) (*proto.DriverResponse, error)
	GetByID(id int32) (*proto.DriverResponse, error)
	GetAll() ([]*proto.DriverResponse, error)
	Update(driver *proto.UpdateDriverRequest) (*proto.DriverResponse, error)
	Delete(id int32) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(driver *proto.CreateDriverRequest) (*proto.DriverResponse, error) {
	return s.repo.Create(driver)
}

func (s *service) GetByID(id int32) (*proto.DriverResponse, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAll() ([]*proto.DriverResponse, error) {
	return s.repo.GetAll()
}

func (s *service) Update(driver *proto.UpdateDriverRequest) (*proto.DriverResponse, error) {
	return s.repo.Update(driver)
}

func (s *service) Delete(id int32) error {
	return s.repo.Delete(id)
}
