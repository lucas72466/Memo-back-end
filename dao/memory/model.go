package memory

// 定义数据库结构

type Comment struct {
	ID            int64  `json:"id" gorm:"column:id; autoIncrement"`
	Author        string `json:"author" gorm:"column:author"`
	Content       string `json:"content" gorm:"column:content"`
	Anonymously   int    `json:"anonymously" gorm:"column:anonymously"`
	PublicVisible int    `json:"public_visible" gorm:"column:public_visible"`
	BuildingID    int64  `json:"building_id" gorm:"column:building_id"`
	CreateTime    int64  `json:"create_time" gorm:"column:create_time; autoCreateTime:milli"`
	UpdateTime    int64  `json:"update_time" gorm:"column:update_time; autoUpdateTime:milli"`
}

func (comment *Comment) TableName() string {
	return "comment"
}
