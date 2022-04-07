package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type DefaultResponse struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

// TODO 加入日志和链路相关逻辑，本次实现最简单版本

func ResponseError(c *gin.Context, resp *DefaultResponse, err error) {
	c.JSON(200, resp)
	LogWithContext(c, ErrorLevel, fmt.Sprintf("err:%v", err), nil)
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, resp *DefaultResponse) {
	c.JSON(200, resp)
}
