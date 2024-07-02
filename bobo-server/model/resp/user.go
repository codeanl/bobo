package resp

import "time"

type UserInfoResp struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Username      string    `json:"username"`
	Nickname      string    `json:"nickname"`
	Email         string    `json:"email"`
	Avatar        string    `json:"avatar"`
	Status        int       `json:"status"`
	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	Role          int       `json:"role"`
	LastLoginTime time.Time `json:"last_login_time"`
}