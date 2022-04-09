package memory

type Comment struct {
	ID          int64  `json:"id" gorm:"column:id; autoIncrement"`
	Author      string `json:"author" gorm:"column:author"`
	Content     string `json:"content" gorm:"column:content"`
	Anonymously int    `json:"anonymously" gorm:"column:anonymously"`
	Visibility  int    `json:"visibility" gorm:"column:visibility"`
	BuildingID  string `json:"building_id" gorm:"column:building_id"`
	CreateTime  int64  `json:"create_time" gorm:"column:create_time; autoCreateTime:milli"`
	UpdateTime  int64  `json:"update_time" gorm:"column:update_time; autoUpdateTime:milli"`
}

func (comment *Comment) TableName() string {
	return "comment"
}

type Story struct {
	ID           int     `json:"id"              gorm:"column:id; autoIncrement"`
	Author       string  `json:"author"          gorm:"column:author"`
	Title        string  `json:"title"           gorm:"column:title"`
	Content      *string `json:"content"         gorm:"column:content"`
	PicturePaths string  `json:"picture_paths"   gorm:"picture_paths"`
	Anonymously  int     `json:"anonymously"     gorm:"anonymously"`
	Visibility   int     `json:"visibility"      gorm:"visibility"`
	BuildingID   string  `json:"building_id"     gorm:"building_id"`
	CreateTime   int64   `json:"create_time"     gorm:"column:create_time; autoCreateTime:milli"`
	UpdateTime   int64   `json:"update_time"     gorm:"column:update_time; autoUpdateTime:milli"`
}

func (story *Story) TableName() string {
	return "story"
}
