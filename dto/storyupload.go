package dto

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
	BuildingID    int    `json:"building_id" binding:"required"`
}

func (param *StoryUploadInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param); err != nil {
		return err
	}
	return nil
}

type StoryUploadOutput struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}
