package model

import "gorm.io/gorm"

type LoginLog struct {
	gorm.Model
	UserId    uint   `gorm:"comment:用户ID" json:"user_id"`
	Type      int    `gorm:"comment:登陆类型（管理段、用户端）" json:"type"`
	IpAddress string `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource  string `gorm:"type:varchar(20);comment:IP" json:"ip_source"`
	Browser   string `gorm:"type:varchar(100);comment:浏览器类型" json:"browser"`
	OS        string `gorm:"type:varchar(100);comment:操作系统" json:"os"`
	Devices   string `gorm:"type:varchar(100);comment:设备信息(如手机、平板、电脑等)" json:"device"`
}
