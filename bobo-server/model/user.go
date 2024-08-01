package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username      string    `gorm:"type:varchar(50);comment:用户名" json:"username"`
	Password      string    `gorm:"type:varchar(100);comment:密码" json:"password"`
	Nickname      string    `gorm:"type:varchar(100);comment:昵称" json:"nickname"`
	Email         string    `gorm:"type:varchar(100);comment:邮箱" json:"email"`
	Avatar        string    `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	Status        int       `gorm:"type:tinyint;comment:状态" json:"status"`
	IpAddress     string    `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string    `gorm:"type:varchar(20);comment:IP" json:"ip_source"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间;null" json:"last_login_time"`
	RoleID        uint      // 外键，指向 Role 表的主键
	Role          Role      `gorm:"foreignkey:RoleID"`
}
