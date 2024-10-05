package app

import "gorm.io/gorm"

// Struct representing Bot object
type Bot struct {
	gorm.Model
	Name      string `gorm:"size:255"`
	Namespace string `gorm:"size:255;uniqueIndex"`
	OwnerID   uint
}
