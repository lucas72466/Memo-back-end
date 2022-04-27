package memory

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type HugAddInput struct {
	MemoryID   int64    `json:"memory_id" binding:"required"`
	MemoryType MemoType `json:"memory_type" binding:"required"`
}

func (param *HugAddInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}
