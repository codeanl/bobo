package front

import (
	"bobo-server/model/req"
	"bobo-server/utils"
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Daily struct{}

// Daily 贴子列表
func (*Daily) DailyList(c *gin.Context) {
	data, count := dailyService.DailyList(utils.BindQuery[req.GetDailyListReq](c))
	r.SuccessListData(c, count, data)
}

// DailySharingInfo 贴子详情
func (*Daily) DailyInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := dailyService.DailyInfo(id)
	r.SendData(c, code, data)
}

// SaveOrUpdate 发布/编辑贴子
func (*Daily) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, dailyService.SaveOrUpdate(c, utils.BindValidJson[req.SaveOrUpdateDailyReq](c), utils.GetFromContext[int](c, "user_id")))
}

// Delete 删除贴子
func (*Daily) Delete(c *gin.Context) {
	r.SendCode(c, dailyService.Delete(utils.BindQuery[req.DeleteDailyReq](c)))
}
