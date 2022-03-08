package controller

import (
	"Memo/conf"
	"Memo/dao/memory"
	"Memo/dto"
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type CommentUploadHandler struct{}

//路由登陆

func CommentUploadRouteRegister(group *gin.RouterGroup) {
	handler := CommentUploadHandler{}
	group.POST("/commentUpload", handler.CommentUpload)
}
func (handler *CommentUploadHandler) CommentUpload(c *gin.Context) {

	//1.绑定校验参数
	param := &dto.CommentUploadInput{}

	if err := param.BindParam(c); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.InvalidParam,
			ErrMsg:  conf.ErrMsg[conf.InvalidParam],
			Data:    nil,
		}, err)
		return
	}

	// 2. 插入数据库

	commentInfo := &memory.CommentInfo{
		Author:        param.Author,
		Content:       param.Content,
		Anonymously:   param.Anonymously,
		PublicVisible: param.PublicVisible,
		BuildingID:    param.BuildingID,
	}

	if err := memory.MDBHandler.CommentUpload(&memory.CommentUploadRequest{
		CommentInfo: commentInfo,
	}); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.InternalError,
			ErrMsg:  conf.ErrMsg[conf.InternalError],
			Data:    nil,
		}, err)

	}

	// 3. 返回状态码和msg
	public.ResponseSuccess(c, &public.DefaultResponse{
		ErrCode: conf.CommentUploadSuccess,
		ErrMsg:  conf.ErrMsg[conf.CommentUploadSuccess],
		Data:    nil,
	})
}
