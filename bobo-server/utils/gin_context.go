package utils

import (
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// Validate 参数合法性校验
func Validate(c *gin.Context, data any) {
	validMsg := Validator.Validate(data)
	if validMsg != "" {
		r.ReturnJson(c, http.StatusOK, r.ERROR_INVALID_PARAM, validMsg, nil)
		panic(nil)
	}
}

// BindJson 绑定
func BindJson[T any](c *gin.Context) (data T) {
	if err := c.ShouldBindJSON(&data); err != nil {
		Logger.Error("BindJson", zap.Error(err))
		panic(r.ERROR_REQUEST_PARAM)
	}
	return
}

// BindValidJson 绑定验证 + 合法性校验
func BindValidJson[T any](c *gin.Context) (data T) {
	// Json 绑定
	if err := c.ShouldBindJSON(&data); err != nil {
		Logger.Error("BindValidJson", zap.Error(err))
		panic(r.ERROR_REQUEST_PARAM)
	}
	// 参数合法性校验
	Validate(c, &data)
	return data
}

// 从 Gin Context 上获取值, 该值是 JWT middleware 解析 Token 后设置的
// 如果该值不存在, 说明 Token 有问题
func GetFromContext[T any](c *gin.Context, key string) T {
	val, exist := c.Get(key)
	if !exist {
		panic(r.ERROR_TOKEN_RUNTIME)
	}
	return val.(T)
}
