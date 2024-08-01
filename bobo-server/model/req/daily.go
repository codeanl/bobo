package req

import "bobo-server/model"

type GetDailyListReq struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	UserId   string `form:"user_id"`
}

type SaveOrUpdateDailyReq struct {
	ID         uint                    `json:"id"`
	Content    string                  `json:"content"`
	Attachment []model.DailyAttachment `json:"attachment"`
	Topic      []string                `json:"topic"`
}

type DeleteDailyReq struct {
	ID int `form:"id"`
}
