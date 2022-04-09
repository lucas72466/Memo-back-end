package memory

import memoryDTO "Memo/dto/memory"

type DBHandler interface {
	CreateStory(req *CreateStoryRequest) error
	CreateComment(req *CreateCommentRequest) error
	SearchComment(req *SearchCommentRequest) (*SearchCommentResult, error)
}

type StoryInfo struct {
	Title       string   `json:"title"`
	Content     *string  `json:"content"`
	PicturePath []string `json:"picture_link"`
	Author      string   `json:"author"`
	Anonymously int      `json:"anonymously"`
	Visibility  int      `json:"visibility"`
	BuildingID  string   `json:"building_id"`
}

type CreateStoryRequest struct {
	StoryInfo *StoryInfo
}

type CommentInfo struct {
	ID          int64                 `json:"ID"`
	Author      string                `json:"author"`
	Content     string                `json:"content"`
	Anonymously memoryDTO.Anonymously `json:"anonymously "`
	Visibility  memoryDTO.Visibility  `json:"visibility"`
	BuildingID  string                `json:"building_id"`
}

type CreateCommentRequest struct {
	CommentInfo *CommentInfo
}

type SearchCommentRequest struct {
	BuildingID string `json:"building_id"`
	Author     string `json:"author"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	PageSize   int    `json:"page_size"`
	Page       int    `json:"page"`
}

type SearchCommentResult struct {
	Comments []*CommentInfo `json:"comments"`
	Total    int32          `json:"total"`
}
