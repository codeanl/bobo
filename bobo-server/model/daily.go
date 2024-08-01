package model

import "gorm.io/gorm"

type Daily struct {
	gorm.Model
	Content         string            `gorm:"type:varchar(255);comment:内容" json:"content"`
	IPLoc           string            `gorm:"type:varchar(255);comment:发布地址" json:"ip_loc"`
	ISTop           int               `gorm:"type:tinyint;comment:置顶" json:"is_top"`
	ISDelete        int               `gorm:"type:tinyint;comment:是否删除" json:"is_delete"`
	UserID          uint              // 外键，指向 User 表的主键
	User            User              `gorm:"foreignkey:UserID"`
	Topic           []*Topic          `gorm:"many2many:daily_topic;joinForeignKey:daily_id" json:"topic"`
	DailyAttachment []DailyAttachment `gorm:"foreignkey:DailyId"`
}

type DailyTopic struct {
	DailyId int
	TopicId int
}
