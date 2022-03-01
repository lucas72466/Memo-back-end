package public

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type DefaultResponse struct {
	ErrCode int         `json:"error_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

// TODO 加入日志和链路相关逻辑，本次实现最简单版本

func ResponseError(c *gin.Context, resp *DefaultResponse, err error) {
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, resp *DefaultResponse) {
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
