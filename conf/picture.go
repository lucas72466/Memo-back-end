package conf

import "time"

const (
	PictureUploadMemoryLimit = 32 << 20
	PictureUploadKey         = "file"
	PictureRelateBuildIDKey  = "building_id"
	PictureStorageBucketName = "memo-backend"
	PictureUploadTimeOut     = time.Second * 5
)
