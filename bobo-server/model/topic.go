package model

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name  string   `gorm:"unique;type:varchar(20);not null;comment:话题;" json:"name"`
	Daily []*Daily `gorm:"many2many:daily_topic;" json:"daily,omitempty"`
	//omitempty 是一个选项，表示如果该字段的值为其类型的零值（例如空字符串、0、false、nil 等），则在序列化时忽略该字段，不将其包含在 JSON 输出中。
}
