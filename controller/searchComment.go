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

type SearchCommentHandler struct {
	c        *gin.Context
	req      *memoryDTO.SearchCommentInput
	respData *memoryDTO.SearchCommentOutputData

	username             string
	login                bool
	searchedCommentInfos []*memoryDAO.CommentInfo
	returnComments       []*memoryDTO.Comment
}

func NewSearchCommentHandler() *SearchCommentHandler {
	handler := &SearchCommentHandler{
		req:      &memoryDTO.SearchCommentInput{},
		respData: &memoryDTO.SearchCommentOutputData{},
	}

	return handler
}

func SearchCommentRouteRegister(group *gin.RouterGroup) {
	group.POST("/comment", NewSearchCommentHandler().SearchComment)
}

func (handler *SearchCommentHandler) SearchComment(c *gin.Context) {
	handler.c = c
	for _, handleFunc := range []func() (conf.StatusCode, error){
		handler.bindParams, handler.checkParams, handler.checkLoginStatus,
		handler.search, handler.filterCommentInfosAndConvert,
	} {
		statusCode, err := handleFunc()
		if err != nil {
			handler.makeResponse(statusCode, err)
			return
		}
	}

	handler.makeResponse(conf.Success, nil)
}

func (handler *SearchCommentHandler) bindParams() (conf.StatusCode, error) {
	if err := handler.req.BindParam(handler.c); err != nil {
		return conf.InvalidParam, err
	}
	return conf.Success, nil
}

func (handler *SearchCommentHandler) checkParams() (conf.StatusCode, error) {
	req := handler.req
	if req == nil || req.Condition == nil {
		return conf.InvalidParam, errors.New("search condition can not be empty")
	}
	condition := req.Condition
	if allEmpty := public.CheckIsStringParamsAllEmpty(condition.BuildingID, condition.Author,
		strconv.FormatInt(condition.StartTime, conf.Decimal), strconv.FormatInt(condition.EndTime, conf.Decimal)); allEmpty {
		return conf.InvalidParam, errors.New("at least one search condition should be set")
	}

	return conf.Success, nil
}

func (handler *SearchCommentHandler) checkLoginStatus() (conf.StatusCode, error) {
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

func (handler *SearchCommentHandler) search() (conf.StatusCode, error) {
	req := handler.req
	condition := req.Condition

	searchRequest := &memoryDAO.SearchCommentRequest{
		BuildingID: condition.BuildingID,
		Author:     condition.Author,
		StartTime:  condition.StartTime,
		EndTime:    condition.EndTime,
		PageSize:   req.PageSize,
		Page:       req.Page,
	}

	result, err := memoryDAO.MDBHandler.SearchComment(searchRequest)
	if err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.InternalError, err
	}

	handler.searchedCommentInfos = result.Comments

	return conf.Success, nil
}

func (handler *SearchCommentHandler) filterCommentInfosAndConvert() (conf.StatusCode, error) {
	res := make([]*memoryDTO.Comment, 0, len(handler.searchedCommentInfos))
	for _, commentInfo := range handler.searchedCommentInfos {
		if !handler._shouldBeReturn(commentInfo) {
			continue
		}

		res = append(res, handler._convertSingleCommentInfo2DTOComment(commentInfo))
	}

	handler.returnComments = res

	return conf.Success, nil
}

func (handler *SearchCommentHandler) makeResponse(statusCode conf.StatusCode, err error) {
	resp := &public.DefaultResponse{
		StatusCode: statusCode.Code,
		Msg:        statusCode.Msg,
	}

	if err != nil {
		public.ResponseError(handler.c, resp, err)
		return
	}

	handler.respData.Count = len(handler.returnComments)
	handler.respData.Comments = handler.returnComments

	resp.Data = handler.respData

	public.ResponseSuccess(handler.c, resp)
}

func (handler *SearchCommentHandler) _shouldBeReturn(commentInfo *memoryDAO.CommentInfo) bool {
	/* Rule:
	   1 .If the current user is not yet logged in， just show the VisibleToAll comment
	   2. If current user has logged in, show all the VisibleToAll comment (2.1)&
	      VisibleToMySelf but author same as username comment (2.2)
	   3. If current user has logged in and he/she try to search other people's comment, just show the VisibleToAll comment (v1 do not need that)
	*/

	// 1.
	if !handler.login {
		return commentInfo.Visibility == memoryDTO.VisibleToAll
	}

	// --- ( has login ) ---
	// 2.1
	if commentInfo.Visibility == memoryDTO.VisibleToAll {
		return true
	}

	// 2.2
	if commentInfo.Visibility == memoryDTO.VisibleToMySelf && commentInfo.Author == handler.username {
		return true
	}

	return false
}

func (handler *SearchCommentHandler) _convertSingleCommentInfo2DTOComment(commentInfo *memoryDAO.CommentInfo) *memoryDTO.Comment {
	/* Rule of Anonymously
	   1. If current user is not logged in, commentInfo which is AnonymouslyYes will always hide true author (replace with other name)
	   2. If current user has logged in and author name is same as user, show the author even if AnonymouslyNo
	*/

	author := commentInfo.Author
	if handler.username != author && commentInfo.Anonymously == memoryDTO.AnonymouslyYes {
		author = "***"
	}

	comment := &memoryDTO.Comment{
		CommentID:  commentInfo.ID,
		Author:     author,
		BuildingID: commentInfo.BuildingID,
		Content:    commentInfo.Content,
		CreateTime: commentInfo.CreateTime,
		UpdateTime: commentInfo.UpdateTime,
	}

	return comment
}
