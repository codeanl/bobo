package front

import (
	"bobo-server/model/req"
	"bobo-server/utils"
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
)

type DailySharing struct{}

// DailySharingList 贴子列表
func (*DailySharing) DailySharingList(c *gin.Context) {
	data, count := dailySharingService.DailySharingList(utils.BindQuery[req.GetDailySharingListReq](c))
	r.SuccessListData(c, count, data)
}

// DailySharingInfo 贴子详情
func (*DailySharing) DailySharingInfo(c *gin.Context) {
	data := dailySharingService.DailySharingInfo(utils.BindQuery[req.GetDailySharingInfoReq](c))
	r.SuccessData(c, data)
}

// SaveOrUpdate 发布/编辑贴子
func (*DailySharing) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, dailySharingService.SaveOrUpdate(c, utils.BindValidJson[req.SaveOrUpdateDailySharingReq](c), utils.GetFromContext[int](c, "user_id")))
}

// Delete 删除贴子
func (*DailySharing) Delete(c *gin.Context) {
	r.SendCode(c, dailySharingService.Delete(utils.BindQuery[req.DeleteDailySharingReq](c)))
}
