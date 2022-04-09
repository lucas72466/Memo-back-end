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

func convertPicRelativePathsFromMySQLSingleString(pathsString string) []string {
	res := strings.Split(pathsString, DefaultPathSplitSymbol)
	res = res[:len(res) - 1]
	return res
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
		CreateTime:  comment.CreateTime,
		UpdateTime:  comment.UpdateTime,
	}

	return commentInfo
}

func convertDBStories2StoryInfos(stories []*Story) []*StoryInfo {
	res := make([]*StoryInfo, len(stories))
	for idx, story := range stories {
		res[idx] = convertSingleStory2StoryInfo(story)
	}

	return res
}

func convertSingleStory2StoryInfo(story *Story) *StoryInfo {
	storyInfo := &StoryInfo{
		ID:           story.ID,
		Title:        story.Title,
		Content:      story.Content,
		PicturePaths: convertPicRelativePathsFromMySQLSingleString(story.PicturePaths),
		Author:       story.Author,
		Anonymously:  memoryDTO.Anonymously(story.Anonymously),
		Visibility:   memoryDTO.Visibility(story.Visibility),
		BuildingID:   story.BuildingID,
		CreateTime:   story.CreateTime,
		UpdateTime:   story.UpdateTime,
	}

	return storyInfo
}
