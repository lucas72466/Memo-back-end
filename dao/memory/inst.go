package memory

import (
	"errors"
	"gorm.io/gorm"
)

type MySQLDBHandler struct {
	MySQLInst *gorm.DB
}

func (handler *MySQLDBHandler) CommentUpload(req *CommentUploadRequest) error {
	if req == nil || req.CommentInfo == nil {
		return errors.New("comment info can not be empty")
	}

	comment := &Comment{
		Author:        req.CommentInfo.Author,
		Content:       req.CommentInfo.Content,
		Anonymously:   req.CommentInfo.Anonymously,
		PublicVisible: req.CommentInfo.PublicVisible,
		BuildingID:    req.CommentInfo.BuildingID,
	}

	if err := handler.MySQLInst.Debug().Create(&comment).Error; err != nil {
		return err
	}

	return nil
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
