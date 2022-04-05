package controller

import (
	"github.com/gin-gonic/gin"
)

type StoryUploadHandler struct {
}

func StoryUploadRouteRegister(group *gin.RouterGroup) {
	handler := StoryUploadHandler{}
	group.POST("/storyUpload", handler.StoryUpload)
}

func (handler *StoryUploadHandler) StoryUpload(c *gin.Context) {

	//// 1. 绑定
	//// 1.1校验参数
	//param := &dto.StoryUploadInput{}
	//
	//if err := param.BindParam(c); err != nil {
	//	public.ResponseError(c, &public.DefaultResponse{
	//		ErrCode: conf.InvalidParam,
	//		ErrMsg:  conf.ErrMsg[conf.InvalidParam],
	//		Data:    nil,
	//	}, err)
	//	return
	//}
	//
	//// 1.2若content和picture皆空，则返回err
	//if comp := IsEmptyContAndPic(param.Content, param.PictureLink); comp == true {
	//	public.ResponseError(c, &public.DefaultResponse{
	//		ErrCode: conf.EmptyContentAndPicture,
	//		ErrMsg:  conf.ErrMsg[conf.EmptyContentAndPicture],
	//		Data:    nil,
	//	}, errors.New(""))
	//	return
	//}
	//
	//// token中获取username
	//info, err := public.GetUserTokenInfoFromContext(c)
	//if err != nil {
	//	public.ResponseError(c, &public.DefaultResponse{
	//		ErrCode: conf.InternalError,
	//		ErrMsg:  conf.ErrMsg[conf.InternalError],
	//		Data:    nil,
	//	}, err)
	//	return
	//}
	//
	//// 2.在数据库中生成记录
	//err = memory.MDBHandler.UploadStory(&memory.StoryUploadRequest{
	//	StoryInfo: &memory.StoryInfo{
	//		Title:         param.Title,
	//		Content:       param.Content,
	//		PictureLink:   param.PictureLink,
	//		Author:        info.UserName,
	//		Anonymously:   param.Anonymously,
	//		PublicVisible: param.PublicVisible,
	//		BuildingID:    param.BuildingID,
	//	}})
	//if err != nil {
	//	public.ResponseError(c, &public.DefaultResponse{
	//		ErrCode: conf.InternalError,
	//		ErrMsg:  conf.ErrMsg[conf.InternalError],
	//		Data:    nil,
	//	}, err)
	//	return
	//}
	//
	//// 3.返回状态
	//public.ResponseSuccess(c, &public.DefaultResponse{
	//	ErrCode: conf.StoryUploadSuccess,
	//	ErrMsg:  conf.ErrMsg[conf.StoryUploadSuccess],
	//	Data:    nil,
	//})

}

// 比较content和picture是否都为空

func IsEmptyContAndPic(content string, picture string) bool {
	if comp := IsEmpty(content) && IsEmpty(picture); comp == true {
		return true
	}

	return false
}

func IsEmpty(param string) bool {
	if param == "" {
		return true
	}

	return false
}
