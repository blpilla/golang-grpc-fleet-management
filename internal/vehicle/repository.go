package vehicle

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(vehicle *Vehicle) error
	GetByID(id uint) (*Vehicle, error)
	GetAll() ([]Vehicle, error)
	Update(vehicle *Vehicle) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(vehicle *Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *repository) GetByID(id uint) (*Vehicle, error) {
	var vehicle Vehicle
	if err := r.db.First(&vehicle, id).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *repository) GetAll() ([]Vehicle, error) {
	var vehicles []Vehicle
	if err := r.db.Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (r *repository) Update(vehicle *Vehicle) error {
	return r.db.Save(vehicle).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Vehicle{}, id).Error
}
