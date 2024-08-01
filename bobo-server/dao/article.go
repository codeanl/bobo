package dao

import (
	"bobo-server/model/req"
	"bobo-server/model/resp"
)

type Article struct{}

func (*Article) GetArticleList(req req.GetArticleListReq) ([]resp.GetArticleListResp, int64) {
	var resp []resp.GetArticleListResp
	var total int64
	db := DB.Table("article").Select("*").Count(&total).Preload("Category").Find(&resp)
	if req.UserId != "" {
		db = db.Where("user_id = ?", req.UserId)
	}
	if req.Title != "" {
		db = db.Where("title like ?", req.Title)
	}
	if req.CategoryId != 0 {
		db = db.Where("category_id = ?", req.CategoryId)
	}
	if req.Type != 0 {
		db = db.Where("type = ?", req.Type)
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}
	if req.IsDelete != 0 {
		db = db.Where("is_delete = ?", req.IsDelete)
	}
	if req.IsReview != 0 {
		db = db.Where("is_review = ?", req.IsReview)
	}
	if req.PageNum > 0 && req.PageSize > 0 {
		db.Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1))
	}
	return resp, total
}

func (d *Article) GetArticle(id int) (*resp.GetArticleListResp, error) {
	var resp resp.GetArticleListResp
	if err := DB.Table("article").Select("*").Preload("Category").Where("id=?", id).First(&resp).Error; err != nil {
		return nil, err
	}
	return &resp, nil
}
