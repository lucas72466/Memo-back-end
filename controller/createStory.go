package controller

import (
	"Memo/conf"
	memoryDAO "Memo/dao/memory"
	memoryDTO "Memo/dto/memory"
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type CreateStoryHandler struct {
	c   *gin.Context
	req *memoryDTO.CreateStoryInput

	username             string
	pictureRelativePaths []string
}

func NewCreateStoryHandler() *CreateStoryHandler {
	return &CreateStoryHandler{
		req: &memoryDTO.CreateStoryInput{},
	}
}

func CreateStoryRouteRegister(group *gin.RouterGroup) {
	group.POST("story/create", NewCreateStoryHandler().CreateStory)
}

func (handler *CreateStoryHandler) CreateStory(c *gin.Context) {
	handler.c = c
	for _, handleFunc := range []func() (conf.StatusCode, error){
		handler.bindParams, handler.getUserInfo, handler.parsePictureRelativePath, handler.create,
	} {
		statusCode, err := handleFunc()
		if err != nil {
			handler.makeResponse(statusCode, err)
			return
		}
	}

	handler.makeResponse(conf.CreateCommentSuccess, nil)
}

func (handler *CreateStoryHandler) bindParams() (conf.StatusCode, error) {
	if err := handler.req.BindParam(handler.c); err != nil {
		return conf.InvalidParam, err
	}
	return conf.Success, nil
}

func (handler *CreateStoryHandler) getUserInfo() (conf.StatusCode, error) {
	info, err := public.GetUserTokenInfoFromContext(handler.c)
	if err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.AuthenticationFail, err
	}

	handler.username = info.UserName

	return conf.Success, nil
}

func (handler *CreateStoryHandler) parsePictureRelativePath() (conf.StatusCode, error) {
	paths := make([]string, len(handler.req.PictureLink))
	for idx, picLink := range handler.req.PictureLink {
		relativePath := public.GetPictureRelativePathFromLink(picLink)
		paths[idx] = relativePath
	}

	handler.pictureRelativePaths = paths

	return conf.Success, nil
}

func (handler *CreateStoryHandler) create() (conf.StatusCode, error) {
	req := handler.req

	story := &memoryDAO.StoryInfo{
		Title:         req.Title,
		Content:       req.Content,
		PicturePath:   handler.pictureRelativePaths,
		Author:        handler.username,
		Anonymously:   int(req.Anonymously),
		PublicVisible: int(req.PublicVisible),
		BuildingID:    req.BuildingID,
	}

	err := memoryDAO.MDBHandler.CreateStory(&memoryDAO.CreateStoryRequest{StoryInfo: story})
	if err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.InternalError, err
	}

	return conf.Success, nil
}

func (handler *CreateStoryHandler) makeResponse(statusCode conf.StatusCode, err error) {
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
