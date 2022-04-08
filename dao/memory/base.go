package memory

type DBHandler interface {
	UploadStory(req *StoryUploadRequest) error
	CreateComment(req *CreateCommentRequest) error
}

type StoryInfo struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	PictureLink   string `json:"picture_link"`
	Author        string `json:"author"`
	Anonymously   bool   `json:"anonymously"`
	PublicVisible int    `json:"publish_visible"`
	BuildingID    string `json:"building_id"`
}

type StoryUploadRequest struct {
	StoryInfo *StoryInfo
}

type CommentInfo struct {
	Author        string `json:"author"`
	Content       string `json:"content"`
	Anonymously   int    `json:"anonymously "`
	PublicVisible int    `json:"publicVisible"`
	BuildingID    string `json:"buildingID"`
}

type CreateCommentRequest struct {
	CommentInfo *CommentInfo
}
