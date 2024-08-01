package model

type DailyAttachment struct {
	Type    string `gorm:"type:varchar(1);comment:类型" json:"type"`
	Url     string `gorm:"type:varchar(255);comment:链接地址" json:"url"`
	DailyId uint   `json:"daily_id"`
}
