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

// Update [单行]更新: 传入对应结构体[传递主键用] 和 带有对应更新字段值的[结构体]，结构体不更新零值
func Update[T any](data *T, slt ...string) {
	if len(slt) > 0 {
		DB.Model(&data).Select(slt).Updates(&data)
		return
	}
	err := DB.Model(&data).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}
