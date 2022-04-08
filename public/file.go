package public

import (
	"Memo/conf"
	"path"
	"strings"
)

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetEncodedFileName(name string) string {
	ext := GetFileExt(name)
	pureFileName := strings.TrimSuffix(name, ext)
	pureFileName = EncodeMD5(pureFileName)

	return pureFileName + ext
}

func GetPictureRelativePathFromLink(link string) string {
	components := strings.Split(link, conf.PictureStorageBucketName)
	return components[len(components)-1]
}
