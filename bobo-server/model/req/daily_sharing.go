package req

type GetDailySharingListReq struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Uid      string `form:"uid"`
}

type GetDailySharingInfoReq struct {
	ID int `form:"id"`
}
