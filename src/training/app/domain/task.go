package domain

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Text   string
	Status string
}
