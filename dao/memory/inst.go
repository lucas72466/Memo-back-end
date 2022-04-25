package memory

import (
	"Memo/conf"
	MemoryDTO "Memo/dto/memory"
	"Memo/public"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
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
		PicturePaths: convertPicRelativePathsToMySQLSingleString(info.PicturePaths),
		Anonymously:  int(info.Anonymously),
		Visibility:   int(info.Visibility),
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
		return nil, errors.New("search request can not be empty")
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

	log.Println(offset, " ", req.Page, "-------------------")

	// search by mysql
	// "zero val" in struct&map condition will not be considered
	// TODO : 1. consider how to design visible and anonymously to reduce return size  2. get total number 3. support time interval filter
	searchedComments := make([]*Comment, 0)
	err := handler.MySQLInst.Debug().Where(&Comment{
		Author:     req.Author,
		BuildingID: req.BuildingID,
	}).Offset(offset).Limit(pageSize).Find(&searchedComments).Error
	if err != nil {
		return nil, fmt.Errorf("search comment fail, err:%w", err)
	}

	return &SearchCommentResult{
		Comments: convertDBComments2CommentInfos(searchedComments),
		Total:    0,
	}, nil
}

func (handler *MySQLDBHandler) SearchStory(req *SearchStoryRequest) (*SearchStoryResult, error) {
	// limit page size to a reasonable value & calculate offset
	pageSize := req.PageSize
	if pageSize <= 0 || pageSize > conf.DefaultCommentPageSizeLimit {
		pageSize = conf.DefaultCommentPageSize
	}
	offset := (req.Page - 1) * pageSize

	searchedStories := make([]*Story, 0)
	err := handler.MySQLInst.Debug().Where(&Story{
		Author:     req.Author,
		Title:      req.Title,
		BuildingID: req.BuildingID,
	}).Offset(offset).Limit(pageSize).Find(&searchedStories).Error
	if err != nil {
		return nil, fmt.Errorf("search story fail, err:%w", err)
	}

	return &SearchStoryResult{
		Stories: convertDBStories2StoryInfos(searchedStories),
		Total:   0,
	}, nil
}

func (handler *MySQLDBHandler) DeleteMemory(req *DeleteMemoryRequest) error {
	if req == nil {
		return errors.New("delete info can not be empty")
	}

	memoryType := req.Type
	var err error
	inst := handler.MySQLInst.Debug()
	switch memoryType {
	case MemoryDTO.MemoTypeComment:
		err = inst.Delete(&Comment{ID: req.MemoryID, Author: req.Author}).Error
	case MemoryDTO.MemoTypeStory:
		err = inst.Delete(&Story{ID: req.MemoryID, Author: req.Author}).Error
	default:
		return fmt.Errorf("unsupport memory type:%v", memoryType)
	}

	return err
}
