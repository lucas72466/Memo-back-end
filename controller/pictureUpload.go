package controller

import (
	"Memo/conf"
	fileDAO "Memo/dao/file"
	fileDTO "Memo/dto/file"
	"Memo/public"
	"context"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type PictureUploadHandler struct {
	c   *gin.Context
	req *fileDTO.PictureUploadInput

	fileHandles  map[string]multipart.File
	locationURLs []string
}

func NewPictureUploadHandler() *PictureUploadHandler {
	return &PictureUploadHandler{req: &fileDTO.PictureUploadInput{}}
}

func PictureUploadRouteRegister(group *gin.RouterGroup) {
	group.POST("/pictureUpload", NewPictureUploadHandler().UploadPicture)
}

func (handler *PictureUploadHandler) UploadPicture(c *gin.Context) {
	handler.c = c
	for _, handleFunc := range []func() (conf.StatusCode, error){
		handler.bindParam, handler.generateUploadFile,
		handler.uploadPicture, handler.releaseFileSource,
	} {
		statusCode, err := handleFunc()
		if err != nil {
			handler.makeResponse(statusCode, err)
			return
		}
	}

	handler.makeResponse(conf.PictureUploadSuccess, nil)
}

func (handler *PictureUploadHandler) bindParam() (conf.StatusCode, error) {
	if err := handler.req.BindParam(handler.c); err != nil {
		public.LogWithContext(handler.c, public.ErrorLevel, err, nil)
		return conf.InvalidParam, err
	}
	public.LogWithContext(handler.c, public.InfoLevel, handler.req, nil)
	return conf.Success, nil
}

func (handler *PictureUploadHandler) generateUploadFile() (conf.StatusCode, error) {
	// get file handle from gin context request
	files := make(map[string]multipart.File, len(handler.req.PicFileHeaders))
	for _, header := range handler.req.PicFileHeaders {
		actualFile, err := header.Open()
		files[header.Filename] = actualFile
		if err != nil {
			public.LogWithContext(handler.c, public.ErrorLevel, err, map[string]interface{}{"filename": header.Filename})
			return conf.InternalError, err
		}
	}

	handler.fileHandles = files

	return conf.Success, nil
}

func (handler *PictureUploadHandler) uploadPicture() (conf.StatusCode, error) {
	c := handler.c

	pictures := make([]*fileDAO.Picture, len(handler.req.PicFileHeaders))
	for idx, header := range handler.req.PicFileHeaders {
		pictures[idx] = &fileDAO.Picture{
			FileHeader: header,
			PicContent: handler.fileHandles[header.Filename],
			SaveName:   public.GetEncodedFileName(header.Filename),
		}
	}

	ctx, _ := context.WithTimeout(c.Request.Context(), conf.PictureUploadTimeOut)
	resp, err := fileDAO.StorageHandler.SavePicture(ctx, &fileDAO.SavePicRequest{
		Pictures:   pictures,
		FolderName: handler.req.BuildingID,
	})
	if err != nil {
		public.LogWithContext(c, public.ErrorLevel, err, nil)
		return conf.InternalError, err
	}

	handler.locationURLs = parseLocationURLsFromLocationMap(resp.Locations)

	return conf.Success, nil
}

func (handler *PictureUploadHandler) releaseFileSource() (conf.StatusCode, error) {
	for fileName, fileHandle := range handler.fileHandles {
		if err := fileHandle.Close(); err != nil {
			public.LogWithContext(handler.c, public.ErrorLevel, err, map[string]interface{}{"filename": fileName})
		}
	}

	return conf.Success, nil
}

func (handler *PictureUploadHandler) makeResponse(statusCode conf.StatusCode, err error) {
	resp := &public.DefaultResponse{
		StatusCode: statusCode.Code,
		Msg:        statusCode.Msg,
	}

	if err != nil {
		public.ResponseError(handler.c, resp, err)
		return
	}

	resp.Data = &fileDTO.PictureUploadOutput{PictureLocationURLs: handler.locationURLs}

	public.ResponseSuccess(handler.c, resp)
}

func parseLocationURLsFromLocationMap(locationMap map[string]string) []string {
	res := make([]string, 0, len(locationMap))
	for _, locationUrl := range locationMap {
		res = append(res, locationUrl)
	}

	return res
}
