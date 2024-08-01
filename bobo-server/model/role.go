package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50);comment:角色名;unique" json:"name"`
	Status int    `gorm:"type:tinyint;comment:状态" json:"status"`
	Label  string `gorm:"type:varchar(50);comment:标识;unique" json:"label"`
	Users  []User `gorm:"foreignKey:RoleID"`
}
