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

type SearchStoryCondition struct {
	BuildingID string `json:"building_id"`
	Author     string `json:"author"`
	Title      string `json:"title"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
}

type SearchStoryInput struct {
	Condition *SearchStoryCondition `json:"condition" binding:"required"`
	PageSize  int                   `json:"page_size" binding:"required"`
	Page      int                   `json:"page" binding:"required"`
}

func (param *SearchStoryInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}

type Story struct {
	ID           int64    `json:"id"`
	Author       string   `json:"author"`
	BuildingID   string   `json:"building_id"`
	Title        string   `json:"title"`
	Content      *string  `json:"content"`
	PictureLinks []string `json:"picture_links"`
	CreateTime   int64    `json:"create_time"`
	UpdateTime   int64    `json:"update_time"`
	HugCount     int      `json:"hug_count"`
	HugStatus    string   `json:"hug_status"`
}

type SearchStoryOutputData struct {
	Stories []*Story `json:"stories"`
	Count   int      `json:"count"`
	Total   int32    `json:"total"`
}
