package dto

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type Anonymously int
type PublicVisible int

var (
	NotAnonymously  Anonymously   = 0
	SetAnonymously  Anonymously   = 1
	VisibleToMyself PublicVisible = 0
	VisibleToFriend PublicVisible = 1
	VisibleToAll    PublicVisible = 2
)

// 定义comment input结构体
type CommentUploadInput struct {
	Content       string `json:"content" binding:"required,max=50,min=1"`
	Anonymously   int    `json:"anonymously"`
	PublicVisible int    `json:"public_visible"`
	BuildingID    int64  `json:"building_id" binding:"required"`
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
