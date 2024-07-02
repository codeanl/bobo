package dao

import (
	"bobo-server/model/req"
	"bobo-server/model/resp"
)

type DailySharing struct{}

func (*DailySharing) GetDailySharingList(req req.GetDailySharingListReq) ([]resp.GetDailySharingListResp, int64) {
	list := make([]resp.GetDailySharingListResp, 0)
	var total int64
	db := DB.Table("daily_sharing").Select("*").Count(&total)
	if req.Uid != "" {
		db = db.Where("uid = ?", req.Uid)
	}
	if req.PageNum > 0 && req.PageSize > 0 {
		db.Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1))
	}
	db.Order("created_at DESC").Find(&list)
	//.Preload("ArticleAttachment")
	return list, total
}
