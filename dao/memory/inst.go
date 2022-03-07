package memory

import (
	"errors"
	"gorm.io/gorm"
)

// TODO 复制粘贴的很乱 要改

type MySQLDBHandler struct {
	MySQLInst *gorm.DB
}

func (handler *MySQLDBHandler) UploadStory(req *StoryUploadRequest) error {
	// TODO 无从下手 我不会了 交给你了orz
	info := req.StoryInfo
	if info == nil {
		return errors.New("story info can not be empty")
	}

	story := Story{
		Author:        info.Author,
		Title:         info.Title,
		Content:       info.Content,
		PictureLink:   info.PictureLink,
		Anonymously:   info.Anonymously,
		PublicVisible: info.PublicVisible,
		BuildingID:    info.BuildingID,
	}
	err := handler.MySQLInst.Debug().Create(&story).Error
	if err != nil {
		return err
	}

	return nil
}
