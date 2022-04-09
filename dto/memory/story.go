package memory

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type CreateStoryInput struct {
	Title       string      `json:"title" binding:"required,max=100" customize_err_msg:"length of title should between 1-100"`
	Content     *string     `json:"content" binding:"required" customize_err_msg:"content can not be empty"`
	PictureLink []string    `json:"picture_link"`
	Anonymously Anonymously `json:"anonymously"`
	Visibility  Visibility  `json:"visibility"`
	BuildingID  string      `json:"building_id" binding:"required"`
}

func (param *CreateStoryInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}
