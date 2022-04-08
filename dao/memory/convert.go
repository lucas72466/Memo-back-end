package memory

import "strings"

const (
	DefaultPathSplitSymbol = ","
)

func convertPicRelativePathsToMySQLSingleString(paths []string) string {
	bf := strings.Builder{}
	for _, path := range paths {
		bf.WriteString(path)
		bf.WriteString(DefaultPathSplitSymbol)
	}

	return bf.String()
}
