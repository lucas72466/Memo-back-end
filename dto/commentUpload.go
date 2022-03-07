package dto

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

// 定义comment input结构体
type CommentUploadInput struct {
	Content       string `json:"content" binding:"required,max=50,min=1"`
	Author        string `json:"author" binding:"required"`
	Anonymously   int    `json:"anonymously " binding:"required"`
	PublicVisible int    `json:"publicVisible" binding:"required"`
	BuildingID    int64  `json:"buildingID" binding:"required"`
}

// 绑定comment参数方法
func (param *CommentUploadInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param); err != nil {
		return err
	}
	return nil
}

// 定义comment output结构体
type CommentUploadOutput struct {
	StatusCode int         `json:"statusCode"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}
