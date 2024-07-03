package req

import "bobo-server/model"

type GetDailySharingListReq struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Uid      string `form:"uid"`
}

type GetDailySharingInfoReq struct {
	ID int `form:"id"`
}

type SaveOrUpdateDailySharingReq struct {
	ID         uint                           `json:"id"`
	Content    string                         `json:"content"`
	Attachment []model.DailySharingAttachment `json:"attachment"`
}

type DeleteDailySharingReq struct {
	ID int `form:"id"`
}
