package memory

// 接口

type DBHandler interface {
	CommentUpload(req *CommentUploadRequest) error
}

type CommentInfo struct {
	Author        string `json:"author"`
	Content       string `json:"content"`
	Anonymously   int    `json:"anonymously "`
	PublicVisible int    `json:"publicVisible"`
	BuildingID    int64  `json:"buildingID"`
}

type CommentUploadRequest struct {
	CommentInfo *CommentInfo
}
