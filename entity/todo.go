package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
	// UserID    uint
}
