package memory

import memoryDTO "Memo/dto/memory"

type DBHandler interface {
	CreateStory(req *CreateStoryRequest) error
	CreateComment(req *CreateCommentRequest) error
	SearchComment(req *SearchCommentRequest) (*SearchCommentResult, error)
	SearchStory(req *SearchStoryRequest) (*SearchStoryResult, error)
	DeleteMemory(req *DeleteMemoryRequest) error
	AddHug(req *AddHugRequest) error
	GetMemoriesRelateHugCount(req *GetMemoriesRelateHugCountRequest) (*GetMemoriesRelateHugCountResult, error)
}

type StoryInfo struct {
	ID           int64                 `json:"ID"`
	Title        string                `json:"title"`
	Content      *string               `json:"content"`
	PicturePaths []string              `json:"picture_paths"`
	Author       string                `json:"author"`
	Anonymously  memoryDTO.Anonymously `json:"anonymously"`
	Visibility   memoryDTO.Visibility  `json:"visibility"`
	BuildingID   string                `json:"building_id"`
	CreateTime   int64                 `json:"create_time"`
	UpdateTime   int64                 `json:"update_time"`
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
	CreateTime  int64                 `json:"create_time"`
	UpdateTime  int64                 `json:"update_time"`
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

type SearchStoryRequest struct {
	BuildingID string `json:"building_id"`
	Author     string `json:"author"`
	Title      string `json:"title"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	PageSize   int    `json:"page_size"`
	Page       int    `json:"page"`
}

type SearchStoryResult struct {
	Stories []*StoryInfo `json:"stories"`
	Total   int32        `json:"total"`
}

type DeleteMemoryRequest struct {
	MemoryID int64              `json:"memory_id"`
	Author   string             `json:"author"`
	Type     memoryDTO.MemoType `json:"type"`
}

type AddHugRequest struct {
	UserName   string             `json:"user_name"`
	MemoryID   int64              `json:"memory_id"`
	MemoryType memoryDTO.MemoType `json:"memory_type"`
}

var (
	TypeStory   int = 1
	TypeComment int = 2
)

type GetMemoriesRelateHugCountRequest struct {
	MemoryIDs  []int64            `json:"memory_ids"`
	MemoryType memoryDTO.MemoType `json:"memory_type"`
}

type GetMemoriesRelateHugCountResult struct {
	MemoriesHugCount map[int64]int `json:"memories_hug_count"`
}
