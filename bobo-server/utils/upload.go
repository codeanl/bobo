package utils

import (
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) (string, int) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return "", r.FAIL
	}
	// 指定文件保存的路径
	savePath := "./uploads/" + file.Filename
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		return "", r.FAIL
	}
	// 返回成功信息
	data := "http://localhost:4321/uploads/" + file.Filename
	return data, r.OK
}
