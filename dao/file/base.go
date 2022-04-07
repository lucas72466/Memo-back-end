package file

import (
	"context"
	"io"
	"mime/multipart"
)

type FStorageHandler interface {
	SavePicture(ctx context.Context, req *SavePicRequest) (*SavePicResponse, error)
}

type Picture struct {
	FileHeader *multipart.FileHeader
	PicContent io.ReadSeeker
	SaveName   string
}

type SavePicRequest struct {
	FolderName string
	Pictures   []*Picture
}

type SavePicResponse struct {
	Locations map[string]string
}
