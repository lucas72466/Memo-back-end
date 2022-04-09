package controller

import (
	"Memo/conf"
	memoryDAO "Memo/dao/memory"
	memoryDTO "Memo/dto/memory"
	"Memo/public"
	"errors"
	"github.com/gin-gonic/gin"
)

type CreateCommentHandler struct {
	c   *gin.Context
	req *memoryDTO.CreateCommentInput

	username string
}

func NewCreateCommentHandler() *CreateCommentHandler {
	return &CreateCommentHandler{
		req: &memoryDTO.CreateCommentInput{},
	}
}

func CreateCommentRouteRegister(group *gin.RouterGroup) {
	group.POST("/comment", NewCreateCommentHandler().CreateComment)
}

func (handler *CreateCommentHandler) CreateComment(c *gin.Context) {
	handler.c = c
	for _, handleFunc := range []func() (conf.StatusCode, error){
		handler.bindParams, handler.getUserInfo, handler.create,
	} {
		statusCode, err := handleFunc()
		if err != nil {
			handler.makeResponse(statusCode, err)
			return
		}
	}

	handler.makeResponse(conf.CreateCommentSuccess, nil)
}

func (handler *CreateCommentHandler) bindParams() (conf.StatusCode, error) {
	if err := handler.req.BindParam(handler.c); err != nil {
		return conf.InvalidParam, err
	}
	return conf.Success, nil
}

func (handler *CreateCommentHandler) getUserInfo() (conf.StatusCode, error) {
	info, err := public.GetUserTokenInfoFromContext(handler.c)
	if err != nil {
		return conf.AuthenticationFail, err
	}

	handler.username = info.UserName

	return conf.Success, nil
}

func (handler *CreateCommentHandler) create() (conf.StatusCode, error) {
	req := handler.req
	commentInfo := &memoryDAO.CommentInfo{
		Author:      handler.username,
		Content:     req.Content,
		Anonymously: req.Anonymously,
		Visibility:  req.Visibility,
		BuildingID:  req.BuildingID,
	}

	if err := memoryDAO.MDBHandler.CreateComment(&memoryDAO.CreateCommentRequest{CommentInfo: commentInfo}); err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.InternalError, errors.New("save comment to db fail")
	}

	return conf.Success, nil
}

func (handler *CreateCommentHandler) makeResponse(statusCode conf.StatusCode, err error) {
	resp := &public.DefaultResponse{
		StatusCode: statusCode.Code,
		Msg:        statusCode.Msg,
	}

	if err != nil {
		public.ResponseError(handler.c, resp, err)
		return
	}

	public.ResponseSuccess(handler.c, resp)
}
