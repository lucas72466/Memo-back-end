package controller

import (
	"Memo/conf"
	memoryDAO "Memo/dao/memory"
	memoryDTO "Memo/dto/memory"
	"Memo/public"
	"errors"
	"github.com/gin-gonic/gin"
)

type CommentUploadHandler struct {
	c    *gin.Context
	req  *memoryDTO.CommentUploadInput

	username string
}

func NewCommentUploadHandler() *CommentUploadHandler {
	return &CommentUploadHandler{
		req:  &memoryDTO.CommentUploadInput{},
	}
}

func CommentUploadRouteRegister(group *gin.RouterGroup) {
	group.POST("/commentUpload", NewCommentUploadHandler().UploadComment)
}

func (handler *CommentUploadHandler) UploadComment(c *gin.Context) {
	handler.c = c
	for _, handleFunc := range []func()(conf.StatusCode, error) {
		handler.bindParams, handler.getUserInfo, handler.upload,
	} {
		statusCode, err := handleFunc()
		if err != nil {
			handler.makeResponse(statusCode, err)
			return
		}
	}

	handler.makeResponse(conf.CommentUploadSuccess, nil)
}

func (handler *CommentUploadHandler) bindParams() (conf.StatusCode, error) {
	if err := handler.req.BindParam(handler.c); err != nil {
		return conf.InvalidParam, err
	}
	return conf.Empty, nil
}

func (handler *CommentUploadHandler) getUserInfo() (conf.StatusCode, error) {
	info, err := public.GetUserTokenInfoFromContext(handler.c)
	if err != nil {
		return conf.AuthenticationFail, err
	}

	handler.username = info.UserName

	return conf.Empty, nil
}

func (handler *CommentUploadHandler) upload() (conf.StatusCode, error) {
	req := handler.req
	commentInfo := &memoryDAO.CommentInfo{
		Author:        handler.username,
		Content:       req.Content,
		Anonymously:   req.Anonymously,
		PublicVisible: req.PublicVisible,
		BuildingID:    req.BuildingID,
	}

	if err := memoryDAO.MDBHandler.CommentUpload(&memoryDAO.CommentUploadRequest{CommentInfo: commentInfo}); err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.InternalError, errors.New("save comment to db fail")
	}

	return conf.Empty, nil
}

func (handler *CommentUploadHandler) makeResponse(statusCode conf.StatusCode, err error) {
	resp := &public.DefaultResponse{
		StatusCode: statusCode,
	}

	if err != nil {
		resp.Msg = err.Error()
		public.ResponseError(handler.c, resp, err)
		return
	}

	resp.Msg = conf.StatusMsg[statusCode]
	public.ResponseSuccess(handler.c, resp)
}


