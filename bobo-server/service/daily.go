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

type Daily struct{}

// DailyList 贴子列表
func (*Daily) DailyList(req req.GetDailyListReq) ([]resp.GetDailyListResp, int) {
	list, count := dailyDao.GetDailyList(req, 0)
	for item, _ := range list {
		userInfo := dao.GetOne(model.User{}, "id=?", list[item].UserID)
		userInfo.Role = dao.GetOne(model.Role{}, "id=?", userInfo.RoleID)
		list[item].User = utils.CopyProperties[resp.UserInfoResp](userInfo)
		//TODO  评论数 收藏数 分享数 点赞数
		list[item].ViewCount = utils.Redis.ZScore(KEY_DAILY_VIEW_COUNT, strconv.Itoa(int(list[item].ID)))
		list[item].CommentCount = int(dao.Count(model.Comment{}, "topic_id = ? AND type=?", list[item].ID, 1))
	}
	return list, int(count)
}

// DailyInfo 文章详情
func (*Daily) DailyInfo(id int) (data *resp.GetDailyListResp, code int) {
	data, err := dailyDao.GetDaily(id)
	if err != nil {
		return data, r.FAIL
	}
	userInfo := dao.GetOne(model.User{}, "id=?", data.UserID)
	userInfo.Role = dao.GetOne(model.Role{}, "id=?", userInfo.RoleID)
	data.User = utils.CopyProperties[resp.UserInfoResp](userInfo)
	data.ViewCount = utils.Redis.ZScore(KEY_DAILY_VIEW_COUNT, strconv.Itoa(id))
	// * 目前请求一次就会增加访问量, 即刷新可以刷访问量
	utils.Redis.ZincrBy(KEY_DAILY_VIEW_COUNT, strconv.Itoa(id), 1)
	data.CommentCount = int(dao.Count(model.Comment{}, "topic_id = ? AND type=?", data.ID, 1))
	return data, r.OK
}

// SaveOrUpdate 创建/编辑文章
func (*Daily) SaveOrUpdate(c *gin.Context, req req.SaveOrUpdateDailyReq, userId int) int {
	if req.ID > 0 {
		ds := dao.GetOne(model.Daily{}, "id=?", req.ID)
		ds.Content = req.Content
		dao.Update(&ds)
		//
		dao.Delete(model.DailyAttachment{}, "daily_id=?", req.ID)
		dao.Delete(model.DailyTopic{}, "daily_id=?", req.ID)
		//
		var Attachment []model.DailyAttachment
		for _, i := range req.Attachment {
			Attachment = append(Attachment, model.DailyAttachment{
				DailyId: req.ID,
				Type:    i.Type,
				Url:     i.Url,
			})
		}
		if len(Attachment) > 0 {
			dao.Create(&Attachment)
		}
		for _, i := range req.Topic {
			topic := dao.GetOne(model.Topic{}, "name=?", i)
			if topic.ID == 0 {
				dao.Create(&model.Topic{Name: i})
			}
			dao.Create(&model.DailyTopic{DailyId: int(req.ID), TopicId: int(topic.ID)})
		}
	} else {
		ds := &model.Daily{
			Content: req.Content,
			UserID:  uint(userId),
			IPLoc:   utils.IP.GetIpSourceSimpleIdle(utils.IP.GetIpAddress(c)),
			ISTop:   0,
		}
		dao.Create(&ds)
		var Attachment []model.DailyAttachment
		for _, i := range req.Attachment {
			Attachment = append(Attachment, model.DailyAttachment{
				DailyId: ds.ID,
				Type:    i.Type,
				Url:     i.Url,
			})
		}
		if len(Attachment) > 0 {
			dao.Create(&Attachment)
		}
		for _, i := range req.Topic {
			topic := dao.GetOne(model.Topic{}, "name=?", i)
			if topic.ID == 0 {
				topicAdd := &model.Topic{Name: i}
				dao.Create(topicAdd)
				dao.Create(&model.DailyTopic{DailyId: int(ds.ID), TopicId: int(topicAdd.ID)})
			} else {
				dao.Create(&model.DailyTopic{DailyId: int(ds.ID), TopicId: int(topic.ID)})
			}
		}
	}
	return r.OK
}

// Delete 删除文章
func (*Daily) Delete(req req.DeleteDailyReq) int {
	//修改状态即可
	ds := dao.GetOne(model.Daily{}, "id=?", req.ID)
	if ds.ID == 0 {
		return r.ERROR_DAILYSHARING_NOT_EXIST
	} else {
		ds.ISDelete = 2
		dao.Update(&ds, "is_delete")
	}
	//	dao.Delete(model.Daily{}, "id=?", req.ID)
	//	dao.Delete(model.DailyAttachment{}, "daily_id=?", req.ID)
	//	dao.Delete(model.DailyTopic{}, "daily_id=?", req.ID)
	//}
	return r.OK
}
