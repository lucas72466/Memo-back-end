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
