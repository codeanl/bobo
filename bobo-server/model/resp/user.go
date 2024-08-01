package resp

import (
	"bobo-server/model"
	"time"
)

type UserInfoResp struct {
	ID            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	Username      string     `json:"username"`
	Nickname      string     `json:"nickname"`
	Email         string     `json:"email"`
	Avatar        string     `json:"avatar"`
	Status        int        `json:"status"`
	IpAddress     string     `json:"ip_address"`
	IpSource      string     `json:"ip_source"`
	RoleID        uint       `json:"role_id"`
	Role          model.Role `json:"role" gorm:"foreignKey:RoleID"`
	LastLoginTime time.Time  `json:"last_login_time"`
	//TODO
	//ArticleLikeSet []string `json:"article_like_set"`
	//CommentLikeSet []string `json:"comment_like_set"`
}
