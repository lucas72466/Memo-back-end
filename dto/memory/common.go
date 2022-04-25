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

type MemoType string

const (
	MemoTypeComment MemoType = "comment"
	MemoTypeStory   MemoType = "story"
)

type DeleteMemoryInput struct {
	MemoryID   int64    `json:"memory_id" binding:"required"`
	MemoryType MemoType `json:"memory_type" binding:"required"`
}

func (param *DeleteMemoryInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}
