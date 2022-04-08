package memory

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type Anonymously int
type PublicVisible int

const (
	VisibleToAll PublicVisible = iota
	VisibleToFriend
	VisibleToMySelf
)

const (
	AnonymouslyNot Anonymously = iota
	AnonymouslyYes
)

// 定义comment input结构体
type CreateCommentInput struct {
	Content       string        `json:"content" binding:"required,max=50,min=1" customize_err_msg:"length of content should between 1-50"`
	Anonymously   Anonymously   `json:"anonymously"`
	PublicVisible PublicVisible `json:"public_visible"`
	BuildingID    string        `json:"building_id" binding:"required"`
}

// 绑定comment参数方法
func (param *CreateCommentInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}
