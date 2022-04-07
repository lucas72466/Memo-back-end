package file

import (
	"Memo/conf"
	"Memo/public"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type PictureUploadInput struct {
	BuildingID string `json:"building_id" binding:"required"`
	PicFileHeaders  []*multipart.FileHeader `json:"file_headers"`
}

func (param *PictureUploadInput) BindParam(c *gin.Context) error {
	if err := c.Request.ParseMultipartForm(conf.PictureUploadMemoryLimit); err != nil {
		public.LogWithContext(c, public.ErrorLevel, err, nil)
		return err
	}

	files := c.Request.MultipartForm.File
	values := c.Request.MultipartForm.Value

	id, err := getBuildingID(values)
	if err != nil {
		return err
	}

	param.PicFileHeaders = files[conf.PictureUploadKey]
	param.BuildingID = id

	return nil
}

func getBuildingID(values map[string][]string) (string, error) {
	val, ok := values[conf.PictureRelateBuildIDKey]
	if !ok {
		return "", errors.New("building ID does not exist")
	}

	return val[0], nil
}

type PictureUploadOutput struct {
	PictureLocationURLs []string `json:"picture_location_url"`
}
