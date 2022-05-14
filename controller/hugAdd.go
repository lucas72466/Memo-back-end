package controller

import (
	"Memo/conf"
	memoryDAO "Memo/dao/memory"
	MemoryDTO "Memo/dto/memory"
	"Memo/public"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
)

type HugAddHandler struct {
	c   *gin.Context
	req *MemoryDTO.HugAddInput

	username string
}

func NewHugAddHandler() *HugAddHandler {
	return &HugAddHandler{
		req: &MemoryDTO.HugAddInput{},
	}
}

func HugAddRouteRegister(group *gin.RouterGroup) {
	group.POST("/add", NewHugAddHandler().AddHug)
}

func (handler *HugAddHandler) AddHug(c *gin.Context) {
	handler.c = c
	for _, handleFunc := range []func() (conf.StatusCode, error){
		handler.bindParams, handler.tryGetUserInfo, handler.addHug,
	} {
		statusCode, err := handleFunc()
		if err != nil {
			handler.makeResponse(statusCode, err)
			return
		}
	}

	handler.makeResponse(conf.AddHugSuccess, nil)
}

func (handler *HugAddHandler) bindParams() (conf.StatusCode, error) {
	if err := handler.req.BindParam(handler.c); err != nil {
		return conf.InvalidParam, err
	}
	return conf.Success, nil
}

func (handler *HugAddHandler) tryGetUserInfo() (conf.StatusCode, error) {
	claim, err := public.ParseTokenClaimFromContext(handler.c)
	if err != nil {
		public.LogWithContext(handler.c, public.InfoLevel, "current user do not log in yet", nil)
		handler.username = "lovely stranger"
		return conf.Success, nil
	}

	handler.username = claim.UserName

	return conf.Success, nil
}

func (handler *HugAddHandler) addHug() (conf.StatusCode, error) {
	if handler.username == "lovely stranger" && checkHasHugged(handler.c, handler.req) {
		return conf.HasHugged, fmt.Errorf("hugged")
	}

	daoReq := &memoryDAO.AddHugRequest{
		UserName:   handler.username,
		MemoryID:   handler.req.MemoryID,
		MemoryType: handler.req.MemoryType,
	}

	err := memoryDAO.MDBHandler.AddHug(daoReq)
	if err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		if errors.Is(err, memoryDAO.HuggedErr) {
			return conf.HasHugged, err
		}
		return conf.InternalError, err
	}

	return conf.Success, nil
}

func (handler *HugAddHandler) makeResponse(statusCode conf.StatusCode, err error) {
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

var unLoginUserHugRecordMap = cmap.New()

func checkHasHugged(c *gin.Context, req *MemoryDTO.HugAddInput) bool {
	ip := c.ClientIP()
	agent := c.Request.UserAgent()

	key := fmt.Sprintf("%s-%s:%s-%d", ip, agent, req.MemoryType, req.MemoryID)

	if unLoginUserHugRecordMap.Has(key) {
		return true
	}

	unLoginUserHugRecordMap.Set(key, true)

	return false
}