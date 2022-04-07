package public

import (
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
