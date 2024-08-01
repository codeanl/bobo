package dao

import (
	"bobo-server/model/req"
	"bobo-server/model/resp"
)

type Daily struct{}

func (*Daily) GetDailyList(req req.GetDailyListReq, isAll int) ([]resp.GetDailyListResp, int64) {
	var resp []resp.GetDailyListResp
	var total int64

	db := DB.Table("daily").Select("*").Count(&total).Order("is_top desc").Order("created_at desc").Preload("Topic").Preload("DailyAttachment")
	if req.UserId != "" {
		db = db.Where("user_id = ?", req.UserId)
	}
	if isAll == 0 {
		db = db.Where("is_delete = 0")
	}

	if req.PageNum > 0 && req.PageSize > 0 {
		db.Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1))
	}
	db.Find(&resp)
	return resp, total
}

func (d *Daily) GetDaily(id int) (*resp.GetDailyListResp, error) {
	var resp resp.GetDailyListResp
	if err := DB.Table("daily").Select("*").Preload("Topic").Preload("DailyAttachment").Where("id=?", id).First(&resp).Error; err != nil {
		return nil, err
	}
	return &resp, nil
}
