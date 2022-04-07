package file

import (
	"Memo/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

var StorageHandler FStorageHandler

func InitAWSSession() *session.Session {
	session_ := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(conf.SecretKeyID, conf.SecretKey, ""),
		Region:      aws.String(endpoints.EuWest2RegionID),
	}))

	return session_
}

func InitStorageHandler() {
	StorageHandler = NewAWSS3StorageHandler(InitAWSSession())
}
