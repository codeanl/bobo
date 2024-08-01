package req

type GetArticleListReq struct {
	PageSize   int    `form:"page_size"`
	PageNum    int    `form:"page_num"`
	UserId     string `form:"user_id"`
	Title      string `form:"title"`
	CategoryId int    `form:"category_id"`
	Type       int    `form:"type" validate:"required,min=1,max=3"`
	Status     int    `form:"status" validate:"required,min=1,max=3"`
	IsDelete   int    `form:"is_delete" validate:"required,min=0,max=1"`
	IsReview   int    `form:"is_review" validate:"required,min=0,max=1"`
}

type SaveOrUpdateArticleReq struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	Img        string `json:"img"`
	Type       int    `json:"type"` // 1-公开 2-私密
	Status     int    `json:"status"`
	IsDelete   int    `json:"is_delete"`
	IsReview   int    `json:"is_review"` //是否审核
	CategoryId int    `json:"category_id"`
}
type DeleteArticleReq struct {
	ID int `form:"id"`
}
