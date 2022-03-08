package public

import (
	"Memo/conf"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type DefaultResponse struct {
	ErrCode   conf.ResponseCode `json:"error_code"`
	ErrMsg    string            `json:"err_msg"`
	ErrDetail string            `json:"err_detail"`
	Data      interface{}       `json:"data"`
}

// TODO 加入日志和链路相关逻辑，本次实现最简单版本

func ResponseError(c *gin.Context, resp *DefaultResponse, err error) {
	resp.ErrDetail = err.Error()
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
