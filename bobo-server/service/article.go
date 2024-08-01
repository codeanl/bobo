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

type Article struct{}

// ArticleList 文章列表
func (*Article) ArticleList(req req.GetArticleListReq) ([]resp.GetArticleListResp, int) {
	list, count := articleDao.GetArticleList(req)
	for item, _ := range list {
		userInfo := dao.GetOne(model.User{}, "id=?", list[item].UserId)
		userInfo.Role = dao.GetOne(model.Role{}, "id=?", userInfo.RoleID)
		list[item].User = utils.CopyProperties[resp.UserInfoResp](userInfo)
		//TODO  评论数 收藏数 分享数 点赞数
		list[item].ViewCount = utils.Redis.ZScore(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(int(list[item].ID)))
	}
	return list, int(count)
}

// ArticleInfo 文章详情
func (*Article) ArticleInfo(id int) (data *resp.GetArticleListResp, code int) {
	data, err := articleDao.GetArticle(id)
	if err != nil {
		return data, r.FAIL
	}
	userInfo := dao.GetOne(model.User{}, "id=?", data.UserId)
	userInfo.Role = dao.GetOne(model.Role{}, "id=?", userInfo.RoleID)
	data.User = utils.CopyProperties[resp.UserInfoResp](userInfo)
	data.ViewCount = utils.Redis.ZScore(KEY_DAILY_VIEW_COUNT, strconv.Itoa(id))
	// * 目前请求一次就会增加访问量, 即刷新可以刷访问量
	utils.Redis.ZincrBy(KEY_DAILY_VIEW_COUNT, strconv.Itoa(id), 1)
	return data, r.OK
}

// SaveOrUpdate 创建/编辑文章
func (*Article) SaveOrUpdate(c *gin.Context, req req.SaveOrUpdateArticleReq, userId int) int {
	if req.ID > 0 {
		ds := dao.GetOne(model.Article{}, "id=?", req.ID)
		if ds.ID == 0 {
			return r.FAIL
		} else {
			//TODO
			ds.Title = req.Title
			ds.Desc = req.Desc
			ds.Content = req.Content
			ds.Img = req.Img
			ds.Type = req.Type
			ds.IsDelete = req.IsDelete
			ds.IsReview = req.IsReview
			ds.Status = req.Status
			ds.CategoryId = req.CategoryId
			dao.Update(&ds)
		}
	} else {
		ds := utils.CopyProperties[model.Article](req)
		ds.UserId = userId
		ds.IsDelete = 2
		ds.IsReview = 2
		dao.Create(&ds)
	}
	return r.OK
}

// Delete 删除文章
func (*Article) Delete(req req.DeleteArticleReq) int {
	//修改状态即可
	ds := dao.GetOne(model.Article{}, "id=?", req.ID)
	if ds.ID == 0 {
		return r.ERROR_DAILYSHARING_NOT_EXIST
	} else {
		ds.IsDelete = 2
		dao.Update(&ds, "is_delete")
	}
	//	dao.Delete(model.Daily{}, "id=?", req.ID)
	//	dao.Delete(model.DailyAttachment{}, "daily_id=?", req.ID)
	//	dao.Delete(model.DailyTopic{}, "daily_id=?", req.ID)
	//}
	return r.OK
}
