package app

import (
	"gorm.io/gorm"
)

// Struct representing Chat object
type Chat struct {
	gorm.Model
	Namespace string `gorm:"size:255;uniqueIndex"`
	OwnerID   int64
}
