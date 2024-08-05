package driver

import (
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	FirstName string
	LastName  string
	License   string
}
