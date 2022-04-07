package file

import (
	"Memo/conf"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSS3StorageHandler struct {
	awsSession *session.Session
	uploader   *s3manager.Uploader
}

func NewAWSS3StorageHandler(awsSession *session.Session) *AWSS3StorageHandler {
	handler := &AWSS3StorageHandler{awsSession: awsSession}

	handler.uploader = s3manager.NewUploader(awsSession, func(u *s3manager.Uploader) {

	})

	return handler
}

func (handler *AWSS3StorageHandler) SavePicture(ctx context.Context, req *SavePicRequest) (*SavePicResponse, error) {
	if req == nil || req.Pictures == nil || len(req.Pictures) == 0 {
		return nil, errors.New("picture save req is invalid")
	}

	locations := make(map[string]string)
	for _, picture := range req.Pictures {
		input := &s3manager.UploadInput{
			Bucket: aws.String(conf.PictureStorageBucketName),
			Key:    aws.String(req.FolderName + "/" + picture.SaveName),
			Body:   picture.PicContent,
		}

		output, err := handler.uploader.UploadWithContext(ctx, input)
		if err != nil {
			return nil, fmt.Errorf("save %s fail, err:%v", picture.FileHeader.Filename, err)
		}

		locations[picture.SaveName] = output.Location
	}

	return &SavePicResponse{Locations: locations}, nil
}
