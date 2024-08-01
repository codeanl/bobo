package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title      string    `gorm:"type:varchar(100);not null;comment:标题" json:"title"`
	Desc       string    `json:"desc;comment:描述"`
	Content    string    `json:"content;comment:描述"`
	Img        string    `json:"img;comment:封面图片"`
	Type       int       `gorm:"type:tinyint;comment:类型(1-原创 2-转载 3-翻译)" json:"type"` // 1-原创 2-转载 3-翻译
	Status     int       `gorm:"type:tinyint;comment:状态(1-公开 2-私密)" json:"status"`    // 1-公开 2-私密
	IsTop      int       `json:"is_top"`
	IsDelete   int       `json:"is_delete"`
	IsReview   int       `json:"is_review"` //是否审核
	CategoryId int       `json:"category_id"`
	Category   *Category `gorm:"foreignkey:CategoryId" json:"category"`
	UserId     int       `json:"user_id"`
	User       *User     `gorm:"foreignkey:UserId" json:"user"`
}
