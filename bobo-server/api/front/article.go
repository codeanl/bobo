package front

import (
	"bobo-server/model/req"
	"bobo-server/utils"
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Article struct{}

// ArticleList 文章列表
func (*Article) ArticleList(c *gin.Context) {
	data, count := articleService.ArticleList(utils.BindQuery[req.GetArticleListReq](c))
	r.SuccessListData(c, count, data)
}

// ArticleInfo 文章详情
func (*Article) ArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := articleService.ArticleInfo(id)
	r.SendData(c, code, data)
}

// SaveOrUpdate 发布/编辑文章
func (*Article) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, articleService.SaveOrUpdate(c, utils.BindValidJson[req.SaveOrUpdateArticleReq](c), utils.GetFromContext[int](c, "user_id")))
}

// Delete 删除文章
func (*Article) Delete(c *gin.Context) {
	r.SendCode(c, articleService.Delete(utils.BindQuery[req.DeleteArticleReq](c)))
}
