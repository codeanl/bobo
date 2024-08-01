package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"unique;type:varchar(20);not null" json:"name"`
}
