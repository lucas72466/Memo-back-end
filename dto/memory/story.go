package memory

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type StoryUploadInput struct {
	Title         string `json:"title" binding:"required,max=20"`
	Content       string `json:"content"`
	PictureLink   string `json:"picture_link"`
	Anonymously   bool   `json:"anonymously"`
	PublicVisible int    `json:"publish_visible"`
	BuildingID    string `json:"building_id" binding:"required"`
}

func (param *StoryUploadInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}
