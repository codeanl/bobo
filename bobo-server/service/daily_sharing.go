package service

import (
	"bobo-server/dao"
	"bobo-server/model"
	"bobo-server/model/req"
	"bobo-server/model/resp"
	"bobo-server/utils"
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
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

// SaveOrUpdate 创建/编辑文章
func (*DailySharing) SaveOrUpdate(c *gin.Context, req req.SaveOrUpdateDailySharingReq, userId int) int {
	if req.ID > 0 {
		ds := dao.GetOne(model.DailySharing{}, "id=?", req.ID)
		ds.Content = req.Content
		dao.Update(&ds)
		//
		dao.Delete(model.DailySharingAttachment{}, "a_id=?", req.ID)
		//
		for _, i := range req.Attachment {
			dao.Create(&model.DailySharingAttachment{
				AID:  int(ds.ID),
				Type: i.Type,
				Url:  i.Url,
			})
		}
	} else {
		ds := &model.DailySharing{
			Content: req.Content,
			Uid:     userId,
			IPLoc:   utils.IP.GetIpSourceSimpleIdle(utils.IP.GetIpAddress(c)),
			ISTop:   "0",
		}
		dao.Create(&ds)
		for _, i := range req.Attachment {
			dao.Create(&model.DailySharingAttachment{
				AID:  int(ds.ID),
				Type: i.Type,
				Url:  i.Url,
			})
		}
	}
	return r.OK
}

// Delete 删除文章
func (*DailySharing) Delete(req req.DeleteDailySharingReq) int {
	ds := dao.GetOne(model.DailySharing{}, "id=?", req.ID)
	if ds.ID == 0 {
		return r.ERROR_DAILYSHARING_NOT_EXIST
	} else {
		dao.Delete(model.DailySharing{}, "id=?", req.ID)
		dao.Delete(model.DailySharingAttachment{}, "a_id=?", req.ID)
	}
	return r.OK
}
