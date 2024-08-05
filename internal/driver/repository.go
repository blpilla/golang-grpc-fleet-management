package driver

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(driver *Driver) error
	GetByID(id uint) (*Driver, error)
	GetAll() ([]Driver, error)
	Update(driver *Driver) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(driver *Driver) error {
	return r.db.Create(driver).Error
}

func (r *repository) GetByID(id uint) (*Driver, error) {
	var driver Driver
	if err := r.db.First(&driver, id).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *repository) GetAll() ([]Driver, error) {
	var drivers []Driver
	if err := r.db.Find(&drivers).Error; err != nil {
		return nil, err
	}
	return drivers, nil
}

func (r *repository) Update(driver *Driver) error {
	return r.db.Save(driver).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Driver{}, id).Error
}
