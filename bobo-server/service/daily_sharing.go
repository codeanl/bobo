package service

import (
	"bobo-server/dao"
	"bobo-server/model"
	"bobo-server/model/req"
	"bobo-server/model/resp"
	"bobo-server/utils"
	"strconv"
)

type DailySharing struct{}

// DailySharingList 贴子列表
func (*DailySharing) DailySharingList(req req.GetDailySharingListReq) ([]resp.GetDailySharingListResp, int) {
	list, count := dailySharingDao.GetDailySharingList(req)
	for item, i := range list {
		list[item].ViewCount = utils.Redis.ZScore(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(int(list[item].ID)))
		//TODO  评论数 收藏数 分享数 点赞数
		userInfo := dao.GetOne(model.User{}, "id=?", i.Uid)
		list[item].User = resp.UserInfoResp{
			CreatedAt:     userInfo.CreatedAt,
			ID:            userInfo.ID,
			Username:      userInfo.Username,
			Nickname:      userInfo.Nickname,
			Email:         userInfo.Email,
			Avatar:        userInfo.Avatar,
			Status:        userInfo.Status,
			IpAddress:     userInfo.IpAddress,
			IpSource:      userInfo.IpSource,
			LastLoginTime: userInfo.LastLoginTime,
		}
		DailySharingAttachmentList := dao.List([]model.DailySharingAttachment{}, "*", "", "a_id =?", i.ID)
		for _, DailySharingAttachmen := range DailySharingAttachmentList {
			list[item].DailySharingAttachment = append(list[item].DailySharingAttachment, DailySharingAttachmen)
		}
	}
	return list, int(count)
}

// DailySharingInfo 文章详情
func (*DailySharing) DailySharingInfo(req req.GetDailySharingInfoReq) resp.GetDailySharingListResp {
	var articleVo resp.GetDailySharingListResp
	ds := dao.GetOne(model.DailySharing{}, "id=?", req.ID)
	articleVo = utils.CopyProperties[resp.GetDailySharingListResp](ds)
	articleVo.ViewCount = utils.Redis.ZScore(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(req.ID))
	articleVo.User = utils.CopyProperties[resp.UserInfoResp](dao.GetOne(model.User{}, "id=?", ds.Uid))
	//
	DailySharingAttachmentList := dao.List([]model.DailySharingAttachment{}, "*", "", "a_id =?", req.ID)
	for _, DailySharingAttachmen := range DailySharingAttachmentList {
		articleVo.DailySharingAttachment = append(articleVo.DailySharingAttachment, DailySharingAttachmen)
	}
	// * 目前请求一次就会增加访问量, 即刷新可以刷访问量
	utils.Redis.ZincrBy(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(req.ID), 1)
	return articleVo
}
