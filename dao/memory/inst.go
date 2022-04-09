package memory

import (
	"Memo/conf"
	"Memo/public"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type MySQLDBHandler struct {
	MySQLInst *gorm.DB
}

func (handler *MySQLDBHandler) CreateComment(req *CreateCommentRequest) error {
	if req == nil || req.CommentInfo == nil {
		return errors.New("comment info can not be empty")
	}

	comment := &Comment{
		Author:      req.CommentInfo.Author,
		Content:     req.CommentInfo.Content,
		Anonymously: int(req.CommentInfo.Anonymously),
		Visibility:  int(req.CommentInfo.Visibility),
		BuildingID:  req.CommentInfo.BuildingID,
	}

	if err := handler.MySQLInst.Debug().Create(&comment).Error; err != nil {
		return fmt.Errorf("create comment fail, err:%w", err)
	}

	return nil
}

func (handler *MySQLDBHandler) CreateStory(req *CreateStoryRequest) error {
	if req == nil || req.StoryInfo == nil {
		return errors.New("story info can not be empty")
	}

	info := req.StoryInfo
	story := Story{
		Author:       info.Author,
		Title:        info.Title,
		Content:      info.Content,
		PicturePaths: convertPicRelativePathsToMySQLSingleString(info.PicturePath),
		Anonymously:  info.Anonymously,
		Visibility:   info.Visibility,
		BuildingID:   info.BuildingID,
	}
	err := handler.MySQLInst.Debug().Create(&story).Error
	if err != nil {
		return fmt.Errorf("create story fail, err:%w", err)
	}

	return nil
}

func (handler *MySQLDBHandler) SearchComment(req *SearchCommentRequest) (*SearchCommentResult, error) {
	// check param
	if req == nil {
		return nil, errors.New("search requset can not be empty")
	}
	if allEmpty := public.CheckIsStringParamsAllEmpty(req.BuildingID, req.Author,
		strconv.FormatInt(req.StartTime, conf.Decimal), strconv.FormatInt(req.EndTime, conf.Decimal)); allEmpty {
		return nil, errors.New("at least one search condition should be set")
	}

	// limit page size to a reasonable value & calculate offset
	pageSize := req.PageSize
	if pageSize <= 0 || pageSize > conf.DefaultCommentPageSizeLimit {
		pageSize = conf.DefaultCommentPageSize
	}
	offset := (req.Page - 1) * pageSize

	// search by mysql
	// "zero val" in struct&map condition will not be considered
	// TODO : 1. consider how to design visible and anonymously to reduce return size  2. get total number
	searchedComments := make([]*Comment, 0)
	err := handler.MySQLInst.Debug().Where(&Comment{
		Author:     req.Author,
		BuildingID: req.BuildingID,
		CreateTime: req.StartTime,
		UpdateTime: req.EndTime,
	}).Offset(offset).Limit(req.PageSize).Find(&searchedComments).Error
	if err != nil {
		return nil, fmt.Errorf("search comment fail, err:%w", err)
	}

	return &SearchCommentResult{
		Comments: convertDBComments2CommentInfos(searchedComments),
		Total:    0,
	}, nil
}
