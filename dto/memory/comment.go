package memory

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type Anonymously int
type PublicVisible int

const (
	VisibleToMyself PublicVisible = iota + 1
	VisibleToFriend
	VisibleToAll
)

const (
	AnonymouslyNot Anonymously = iota
	AnonymouslyYes
)

// 定义comment input结构体
type CommentUploadInput struct {
	Content       string `json:"content" binding:"required,max=50,min=1" customize_err_msg:"length of content should between 1-50"`
	Anonymously   int    `json:"anonymously"`
	PublicVisible int    `json:"public_visible" binding:"required"`
	BuildingID    string `json:"building_id" binding:"required"`
}

// 绑定comment参数方法
func (param *CommentUploadInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}
