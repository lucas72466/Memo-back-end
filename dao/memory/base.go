package memory

type StoryDBHandler interface {
	UploadStory(req *StoryUploadRequest) error
}

type StoryInfo struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	PictureLink   string `json:"picture_link"`
	Author        string `json:"author"`
	Anonymously   bool   `json:"anonymously"`
	PublicVisible int    `json:"publish_visible"`
	BuildingID    int    `json:"building_id"`
}

type StoryUploadRequest struct {
	StoryInfo *StoryInfo
}
