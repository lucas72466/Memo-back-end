package memory

import (
	memoryDTO "Memo/dto/memory"
	"strings"
)

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

func convertDBComments2CommentInfos(comments []*Comment) []*CommentInfo {
	res := make([]*CommentInfo, len(comments))
	for idx, comment := range comments {
		res[idx] = convertSingleDBComment2CommentInfo(comment)
	}

	return res
}

func convertSingleDBComment2CommentInfo(comment *Comment) *CommentInfo {
	commentInfo := &CommentInfo{
		ID:          comment.ID,
		Author:      comment.Author,
		Content:     comment.Content,
		Anonymously: memoryDTO.Anonymously(comment.Anonymously),
		Visibility:  memoryDTO.Visibility(comment.Visibility),
		BuildingID:  comment.BuildingID,
	}

	return commentInfo
}
