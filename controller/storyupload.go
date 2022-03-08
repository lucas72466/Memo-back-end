package controller

import (
	"Memo/conf"
	"Memo/dao/memory"
	"Memo/dto"
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type StoryUploadHandler struct {
}

func StoryUploadRouteRegister(group *gin.RouterGroup) {
	handler := StoryUploadHandler{}
	group.POST("/memory", handler.StoryUpload)
}

func (handler *StoryUploadHandler) StoryUpload(c *gin.Context) {

	// 1. 绑定校验参数
	param := &dto.StoryUploadInput{}

	if err := param.BindParam(c); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.InvalidParam,
			ErrMsg:  conf.ErrMsg[conf.InvalidParam],
			Data:    nil,
		}, err)
		return
	}

	// 2.在数据库中生成记录
	err := memory.MDBHandler.UploadStory(&memory.StoryUploadRequest{
		StoryInfo: &memory.StoryInfo{
			Title:         param.Title,
			Content:       param.Content,
			PictureLink:   param.PictureLink,
			Author:        param.Author,
			Anonymously:   param.Anonymously,
			PublicVisible: param.PublicVisible,
			BuildingID:    param.BuildingID,
		}})
	if err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.InternalError,
			ErrMsg:  conf.ErrMsg[conf.InternalError],
			Data:    nil,
		}, err)
	}

	// 3.返回状态
	public.ResponseSuccess(c, &public.DefaultResponse{
		ErrCode: conf.StoryUploadSuccess,
		ErrMsg:  conf.ErrMsg[conf.StoryUploadSuccess],
		Data:    nil,
	})

}
