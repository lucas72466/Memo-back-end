package controller

import (
	"Memo/conf"
	memoryDAO "Memo/dao/memory"
	memoryDTO "Memo/dto/memory"
	"Memo/public"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SearchStoryHandler struct {
	c        *gin.Context
	req      *memoryDTO.SearchStoryInput
	respData *memoryDTO.SearchStoryOutputData

	username           string
	login              bool
	searchedStoryInfos []*memoryDAO.StoryInfo
	retStories         []*memoryDTO.Story
}

func NewSearchStoryHandler() *SearchStoryHandler {
	handler := &SearchStoryHandler{
		req:      &memoryDTO.SearchStoryInput{},
		respData: &memoryDTO.SearchStoryOutputData{},
	}

	return handler
}

func SearchStoryRouteRegister(group *gin.RouterGroup) {
	group.POST("/story", NewSearchStoryHandler().SearchStory)
}

func (handler *SearchStoryHandler) SearchStory(c *gin.Context) {
	handler.c = c
	for _, handleFunc := range []func() (conf.StatusCode, error){
		handler.bindParams, handler.checkParams, handler.checkLoginStatus,
		handler.searchStory, handler.filterStoryInfosAndConvert,
	} {
		statusCode, err := handleFunc()
		if err != nil {
			handler.makeResponse(statusCode, err)
			return
		}
	}

	handler.makeResponse(conf.CreateCommentSuccess, nil)
}

func (handler *SearchStoryHandler) bindParams() (conf.StatusCode, error) {
	if err := handler.req.BindParam(handler.c); err != nil {
		return conf.InvalidParam, err
	}
	return conf.Success, nil
}

func (handler *SearchStoryHandler) checkParams() (conf.StatusCode, error) {
	req := handler.req
	if req == nil || req.Condition == nil {
		return conf.InvalidParam, errors.New("search condition can not be empty")
	}
	condition := req.Condition
	if allEmpty := public.CheckIsStringParamsAllEmpty(condition.BuildingID, condition.Author, condition.Title,
		strconv.FormatInt(condition.StartTime, conf.Decimal), strconv.FormatInt(condition.EndTime, conf.Decimal)); allEmpty {
		return conf.InvalidParam, errors.New("at least one search condition should be set")
	}

	return conf.Success, nil
}

func (handler *SearchStoryHandler) checkLoginStatus() (conf.StatusCode, error) {
	claim, err := public.ParseTokenClaimFromContext(handler.c)
	if err != nil {
		public.LogWithContext(handler.c, public.InfoLevel, "current user do not log in yet", nil)
		handler.login = false
		return conf.Success, nil
	}

	handler.login = true
	handler.username = claim.UserName

	return conf.Success, nil
}

func (handler *SearchStoryHandler) searchStory() (conf.StatusCode, error) {
	req := handler.req
	condition := req.Condition

	searchRequest := &memoryDAO.SearchStoryRequest{
		BuildingID: condition.BuildingID,
		Author:     condition.Author,
		Title:      condition.Title,
		StartTime:  condition.StartTime,
		EndTime:    condition.EndTime,
		PageSize:   req.PageSize,
		Page:       req.Page,
	}

	result, err := memoryDAO.MDBHandler.SearchStory(searchRequest)
	if err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.InternalError, err
	}

	handler.searchedStoryInfos = result.Stories

	return conf.Success, nil
}

func (handler *SearchStoryHandler) filterStoryInfosAndConvert() (conf.StatusCode, error) {
	res := make([]*memoryDTO.Story, 0, len(handler.searchedStoryInfos))
	for _, storyInfo := range handler.searchedStoryInfos {
		if !handler._shouldBeReturn(storyInfo) {
			continue
		}

		res = append(res, handler._convertSingleStoryInfo2DTOStory(storyInfo))
	}

	handler.retStories = res

	return conf.Success, nil
}

func (handler *SearchStoryHandler) makeResponse(statusCode conf.StatusCode, err error) {
	resp := &public.DefaultResponse{
		StatusCode: statusCode.Code,
		Msg:        statusCode.Msg,
	}

	if err != nil {
		public.ResponseError(handler.c, resp, err)
		return
	}

	handler.respData.Count = len(handler.retStories)
	handler.respData.Stories = handler.retStories

	resp.Data = handler.respData

	public.ResponseSuccess(handler.c, resp)
}

func (handler *SearchStoryHandler) _shouldBeReturn(storyInfo *memoryDAO.StoryInfo) bool {
	/* Rule:
	   1 .If the current user is not yet logged in， just show the VisibleToAll storyInfo
	   2. If current user has logged in, show all the VisibleToAll storyInfo (2.1)&
	      VisibleToMySelf but author same as username storyInfo (2.2)
	   3. If current user has logged in and he/she try to search other people's storyInfo, just show the VisibleToAll storyInfo (v1 do not need that)
	*/

	// 1.
	if !handler.login {
		return storyInfo.Visibility == memoryDTO.VisibleToAll
	}

	// --- ( has login ) ---
	// 2.1
	if storyInfo.Visibility == memoryDTO.VisibleToAll {
		return true
	}

	// 2.2
	if storyInfo.Visibility == memoryDTO.VisibleToMySelf && storyInfo.Author == handler.username {
		return true
	}

	return false
}

func (handler *SearchStoryHandler) _convertSingleStoryInfo2DTOStory(storyInfo *memoryDAO.StoryInfo) *memoryDTO.Story {
	author := storyInfo.Author
	if handler.username != author && storyInfo.Anonymously == memoryDTO.AnonymouslyYes {
		author = "***"
	}

	story := &memoryDTO.Story{
		ID:           storyInfo.ID,
		Author:       author,
		BuildingID:   storyInfo.BuildingID,
		Title:        storyInfo.Title,
		Content:      storyInfo.Content,
		PictureLinks: storyInfo.PicturePaths,
		CreateTime:   storyInfo.CreateTime,
		UpdateTime:   storyInfo.UpdateTime,
	}

	return story
}
