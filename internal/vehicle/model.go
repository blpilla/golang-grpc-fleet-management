package vehicle

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Make  string
	Model string
}
