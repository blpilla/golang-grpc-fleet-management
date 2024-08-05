package association

import (
	"gorm.io/gorm"
)

type Repository interface {
	AssociateDriverToVehicle(driverID, vehicleID uint) error
	GetAllAssociations() ([]Association, error)
	GetDriversByVehicle(vehicleID uint) ([]uint, error)
	GetVehiclesByDriver(driverID uint) ([]uint, error)
	DissociateDriverFromVehicle(driverID, vehicleID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) AssociateDriverToVehicle(driverID, vehicleID uint) error {
	return r.db.Create(&Association{DriverID: driverID, VehicleID: vehicleID}).Error
}

func (r *repository) GetAllAssociations() ([]Association, error) {
	var associations []Association
	if err := r.db.Find(&associations).Error; err != nil {
		return nil, err
	}
	return associations, nil
}

func (r *repository) GetDriversByVehicle(vehicleID uint) ([]uint, error) {
	var driverIDs []uint
	if err := r.db.Table("associations").Where("vehicle_id = ?", vehicleID).Pluck("driver_id", &driverIDs).Error; err != nil {
		return nil, err
	}
	return driverIDs, nil
}

func (r *repository) GetVehiclesByDriver(driverID uint) ([]uint, error) {
	var vehicleIDs []uint
	if err := r.db.Table("associations").Where("driver_id = ?", driverID).Pluck("vehicle_id", &vehicleIDs).Error; err != nil {
		return nil, err
	}
	return vehicleIDs, nil
}

func (r *repository) DissociateDriverFromVehicle(driverID, vehicleID uint) error {
	return r.db.Where("driver_id = ? AND vehicle_id = ?", driverID, vehicleID).Delete(&Association{}).Error
}
