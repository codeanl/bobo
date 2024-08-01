package resp

import (
	"bobo-server/model"
	"gorm.io/gorm"
)

type GetArticleListResp struct {
	gorm.Model
	Title           string          `json:"title"`
	Desc            string          `json:"desc"`
	Content         string          `json:"content"`
	Img             string          `json:"img"`
	Type            int             `json:"type"`
	Status          int             `json:"status"`
	IsTop           int             `json:"is_top"`
	IsDelete        int             `json:"is_delete"`
	IsReview        int             `json:"is_review"` //是否审核
	CategoryId      int             `json:"category_id"`
	Category        *model.Category `gorm:"foreignkey:CategoryId" json:"category"`
	UserId          int             `json:"user_id"`
	User            UserInfoResp    `gorm:"foreignkey:UserId" json:"user"`
	CommentCount    int             `json:"comment_count"`    //评论数
	CollectionCount int             `json:"collection_count"` //收藏数
	ShareCount      int             `json:"share_count"`      //分享数
	UpvoteCount     int             `json:"upvote_count"`     //点赞数
	ViewCount       int             `json:"view_count"`       //浏览数
}
