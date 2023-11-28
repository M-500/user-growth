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
// @Date 2023/11/28 9:28
//

type CoinUserDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewCoinUserDao(ctx context.Context) *CoinUserDao {
	return &CoinUserDao{
		db:  db_helper.GetDb(),
		ctx: ctx,
	}
}

func (dao *CoinUserDao) Get(id int) (*models.TbCoinUser, error) {
	data := &models.TbCoinUser{}
	_, err := dao.db.ID(id).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dao *CoinUserDao) GetByUid(uid int) (*models.TbCoinUser, error) {
	data := &models.TbCoinUser{}
	sess := dao.db.Where("`uid` = ?", uid)
	_, err := sess.Get(data)
	return data, err
}

func (dao *CoinUserDao) FindAllPager(page, size int) ([]models.TbCoinUser, int64, error) {
	dataList := make([]models.TbCoinUser, 0)
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 15
	}
	start := (page - 1) * size
	total, err := dao.db.Desc("id").Limit(size, start).FindAndCount(&dataList)
	return dataList, total, err
}

// Insert one row
func (dao *CoinUserDao) Insert(data *models.TbCoinUser) error {
	data.SysCreated = common.Now()
	data.SysUpdated = common.Now()
	_, err := dao.db.Insert(data)
	return err
}

// Update one row
func (dao *CoinUserDao) Update(data *models.TbCoinUser, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	_, err := sess.Update(data)
	return err
}

// Save with Insert and Update
func (dao *CoinUserDao) Save(data *models.TbCoinUser, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	} else {
		return dao.Insert(data)
	}
}
