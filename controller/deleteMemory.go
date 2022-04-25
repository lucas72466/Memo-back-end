package controller

import (
	"Memo/conf"
	memoryDAO "Memo/dao/memory"
	memoryDTO "Memo/dto/memory"
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type DeleteMemoryHandler struct {
	c   *gin.Context
	req *memoryDTO.DeleteMemoryInput

	username string
}

func NewDeleteMemoryHandler() *DeleteMemoryHandler {
	return &DeleteMemoryHandler{
		req: &memoryDTO.DeleteMemoryInput{},
	}
}

func DeleteMemoryRouteRegister(group *gin.RouterGroup) {
	group.POST("/", NewDeleteMemoryHandler().DeleteMemory)
}

func (handler *DeleteMemoryHandler) DeleteMemory(c *gin.Context) {
	handler.c = c
	for _, handleFunc := range []func() (conf.StatusCode, error){
		handler.bindParams, handler.getUserInfo, handler.deleteMemory,
	} {
		statusCode, err := handleFunc()
		if err != nil {
			handler.makeResponse(statusCode, err)
			return
		}
	}

	handler.makeResponse(conf.DeleteMemorySuccess, nil)
}

func (handler *DeleteMemoryHandler) bindParams() (conf.StatusCode, error) {
	if err := handler.req.BindParam(handler.c); err != nil {
		return conf.InvalidParam, err
	}
	return conf.Success, nil
}

func (handler *DeleteMemoryHandler) getUserInfo() (conf.StatusCode, error) {
	info, err := public.GetUserTokenInfoFromContext(handler.c)
	if err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.AuthenticationFail, err
	}

	handler.username = info.UserName

	return conf.Success, nil
}

func (handler *DeleteMemoryHandler) deleteMemory() (conf.StatusCode, error) {
	req := &memoryDAO.DeleteMemoryRequest{
		Author:   handler.username,
		MemoryID: handler.req.MemoryID,
		Type:     handler.req.MemoryType,
	}

	err := memoryDAO.MDBHandler.DeleteMemory(req)
	if err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.InternalError, err
	}

	return conf.Success, nil
}

func (handler *DeleteMemoryHandler) makeResponse(statusCode conf.StatusCode, err error) {
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
