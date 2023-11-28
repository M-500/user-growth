package dao

import (
	"context"
	"user_growth/common"
	"user_growth/db_helper"
	"user_growth/models"
	"xorm.io/xorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/28 9:37
//

type GradeInfoDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewGradeInfoDao(ctx context.Context) *GradeInfoDao {
	return &GradeInfoDao{
		db:  db_helper.GetDb(),
		ctx: ctx,
	}
}

func (dao *GradeInfoDao) Get(id int) (*models.TbGradeInfo, error) {
	data := &models.TbGradeInfo{}
	if _, err := dao.db.ID(id).Get(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (dao *GradeInfoDao) FindALL() ([]models.TbGradeInfo, error) {
	dataList := make([]models.TbGradeInfo, 0)
	err := dao.db.Asc("score").Find(&dataList)
	return dataList, err
}

// Insert one row
func (dao *GradeInfoDao) Insert(data *models.TbGradeInfo) error {
	data.SysCreated = common.Now()
	data.SysUpdated = common.Now()
	_, err := dao.db.Insert(data)
	return err
}

// Update one row
func (dao *GradeInfoDao) Update(data *models.TbGradeInfo, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	_, err := sess.Update(data)
	return err
}

// Save with Insert and Update
func (dao *GradeInfoDao) Save(data *models.TbGradeInfo, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	} else {
		return dao.Insert(data)
	}
}
