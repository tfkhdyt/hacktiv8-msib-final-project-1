package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	title     string
	completed bool
	userId    uint
}
