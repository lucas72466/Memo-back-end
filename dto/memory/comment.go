package memory

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type Anonymously int
type Visibility int

const (
	VisibleToAll Visibility = iota
	VisibleToMySelf
)

const (
	AnonymouslyNot Anonymously = iota
	AnonymouslyYes
)

// 定义comment input结构体
type CreateCommentInput struct {
	Content     string      `json:"content" binding:"required,max=50,min=1" customize_err_msg:"length of content should between 1-50"`
	Anonymously Anonymously `json:"anonymously"`
	Visibility  Visibility  `json:"visibility"`
	BuildingID  string      `json:"building_id" binding:"required"`
}

// 绑定comment参数方法
func (param *CreateCommentInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}

type SearchCommentCondition struct {
	BuildingID string `json:"building_id"`
	Author     string `json:"author"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
}

type SearchCommentInput struct {
	Condition *SearchCommentCondition `json:"condition" binding:"required"`
	PageSize  int                     `json:"page_size" binding:"required"`
	Page      int                     `json:"page" binding:"required"`
}

func (param *SearchCommentInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}

type Comment struct {
	CommentID  int64  `json:"comment_id"`
	Author     string `json:"author"`
	BuildingID string `json:"building_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type SearchCommentOutputData struct {
	Comments []*Comment `json:"comments"`
	Count    int        `json:"count"`
	Total    int32      `json:"total"`
}
