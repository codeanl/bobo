package dao

import (
	"errors"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetOne 单条数据查询
func GetOne[T any](data T, query string, args ...any) T {
	err := DB.Where(query, args...).First(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return data
}

// Create 创建数据
func Create[T any](data *T) {
	err := DB.Create(&data).Error
	if err != nil {
		panic(err)
	}
}
