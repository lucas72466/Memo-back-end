package memory

type Story struct {
	ID            int    `json:"id"              gorm:"column:id"`
	Author        string `json:"author"          gorm:"column:author"`
	Title         string `json:"title"           gorm:"column:title"`
	Content       string `json:"content"         gorm:"column:content"`
	PictureLink   string `json:"picture_link"    gorm:"picture_link"`
	Anonymously   bool   `json:"anonymously"     gorm:"anonymously"`
	PublicVisible int    `json:"public_visible"  gorm:"public_visible"`
	BuildingID    int    `json:"building_id"     gorm:"building_id"`
	CreateTime    int64  `json:"create_time"     gorm:"create_time"`
	UpdateTime    int64  `json:"update_time"     gorm:"update_time"`
}

func (story *Story) TableName() string {
	return "story"
}
