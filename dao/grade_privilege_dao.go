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
// @Date 2023/11/28 9:42
//

type GradePrivilegeDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewGradePrivilegeDao(ctx context.Context) *GradePrivilegeDao {
	return &GradePrivilegeDao{
		db:  db_helper.GetDb(),
		ctx: ctx,
	}
}

func (dao *GradePrivilegeDao) Get(id int) (*models.TbGradePrivilege, error) {
	data := &models.TbGradePrivilege{}
	if _, err := dao.db.ID(id).Get(data); err != nil {
		return nil, err
	}
	return data, nil
}

// FindByGrade Find models by gradeId.
func (dao *GradePrivilegeDao) FindByGrade(gradeId int) ([]models.TbGradePrivilege, error) {
	datalist := []models.TbGradePrivilege{}
	if err := dao.db.Where("`grade_id`=?", gradeId).Find(&datalist); err != nil {
		return nil, err
	} else {
		return datalist, nil
	}
}

// FindAll get all models
func (dao *GradePrivilegeDao) FindAll() ([]models.TbGradePrivilege, error) {
	datalist := make([]models.TbGradePrivilege, 0)
	err := dao.db.Desc("id").Find(&datalist)
	return datalist, err
}

// Insert one row
func (dao *GradePrivilegeDao) Insert(data *models.TbGradePrivilege) error {
	data.SysCreated = common.Now()
	data.SysUpdated = common.Now()
	_, err := dao.db.Insert(data)
	return err
}

// Update one row
func (dao *GradePrivilegeDao) Update(data *models.TbGradePrivilege, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	_, err := sess.Update(data)
	return err
}

// Save with Insert and Update
func (dao *GradePrivilegeDao) Save(data *models.TbGradePrivilege, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	} else {
		return dao.Insert(data)
	}
}
