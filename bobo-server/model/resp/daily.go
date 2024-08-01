package resp

import (
	"bobo-server/model"
	"time"
)

type GetDailyListResp struct {
	ID              uint                    `json:"id"`
	CreatedAt       time.Time               `json:"created_at"`
	Content         string                  `json:"content"`
	IPLoc           string                  `json:"ip_loc"`
	ISTop           int                     `json:"is_top"`
	UserID          int                     `json:"user_id"`
	User            UserInfoResp            `json:"user"  gorm:"foreignKey:UserID"`
	Topic           []model.Topic           `gorm:"many2many:daily_topic;joinForeignKey:DailyId" json:"topic"`
	DailyAttachment []model.DailyAttachment `gorm:"foreignkey:DailyId"`
	CommentCount    int                     `json:"comment_count"`    //评论数
	CollectionCount int                     `json:"collection_count"` //收藏数
	ShareCount      int                     `json:"share_count"`      //分享数
	UpvoteCount     int                     `json:"upvote_count"`     //点赞数
	ViewCount       int                     `json:"view_count"`       //浏览数
}
