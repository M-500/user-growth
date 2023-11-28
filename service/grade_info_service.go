package service

import (
	"context"
	"user_growth/dao"
	"user_growth/models"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/28 9:56
//

type GradeInfoService struct {
	ctx          context.Context
	daoGradeInfo *dao.GradeInfoDao
}

func NewGradeInfoService(ctx context.Context) *GradeInfoService {
	return &GradeInfoService{
		ctx:          ctx,
		daoGradeInfo: dao.NewGradeInfoDao(ctx),
	}
}

// Get model by id.
func (s *GradeInfoService) Get(id int) (*models.TbGradeInfo, error) {
	return s.daoGradeInfo.Get(id)
}

// FindAll get all models
func (s *GradeInfoService) FindAll() ([]models.TbGradeInfo, error) {
	return s.daoGradeInfo.FindAll()
}

// Save with Insert and Update
func (s *GradeInfoService) Save(data *models.TbGradeInfo, musColumns ...string) error {
	return s.daoGradeInfo.Save(data, musColumns...)
}

// NowGrade 当前成长值的等级
func (s *GradeInfoService) NowGrade(score int) (*models.TbGradeInfo, error) {
	datalist, err := s.FindAll()
	if err != nil {
		return nil, err
	}
	var grade *models.TbGradeInfo
	for i, data := range datalist {
		if score >= data.Score {
			grade = &datalist[i]
		} else {
			break
		}
	}
	return grade, nil
}
