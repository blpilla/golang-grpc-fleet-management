package association

import (
	"gorm.io/gorm"
)

type Association struct {
	gorm.Model
	DriverID  uint
	VehicleID uint
}
