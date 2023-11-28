package dao

import (
	"context"
	"user_growth/common"
	"user_growth/db_helper"
	"user_growth/models"
	"xorm.io/xorm"
)

//
// @Description 数据层的封装，每个文件对应一张表的增删改查操作
// @Author 代码小学生王木木
// @Date 2023/11/27 15:26
//

type CoinTaskDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewCoinTaskDao(ctx context.Context) *CoinTaskDao {
	return &CoinTaskDao{
		db:  db_helper.GetDb(),
		ctx: ctx,
	}
}

func (dao *CoinTaskDao) Get(id int) (*models.TbCoinTask, error) {
	data := &models.TbCoinTask{}
	_, err := dao.db.ID(id).Get(data)
	if err != nil {
		return nil, err
	}
	if data == nil || data.Id == 0 {
		return nil, nil
	}
	return data, nil
}

func (dao *CoinTaskDao) GetByTask(task string) (*models.TbCoinTask, error) {
	data := &models.TbCoinTask{}
	if _, err := dao.db.Where("`task`=?", task).Get(data); err != nil {
		return nil, err
	} // 预编译方式，避免直接拼接SQL语句造成SQL注入
	if data == nil || data.Id == 0 {
		return nil, nil
	}
	return data, nil
}

func (dao *CoinTaskDao) FindAll() ([]models.TbCoinTask, error) {
	dataList := make([]models.TbCoinTask, 0)
	err := dao.db.Desc("id").Find(&dataList)
	return dataList, err
}

func (dao *CoinTaskDao) Insert(data *models.TbCoinTask) error {
	data.SysCreated = common.Now()
	data.SysUpdated = common.Now()
	_, err := dao.db.Insert(data)
	return err
}

func (dao *CoinTaskDao) Update(data *models.TbCoinTask, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	data.SysUpdated = common.Now()
	_, err := sess.Update(data)
	return err
}

// 更高一层的封装
func (dao *CoinTaskDao) Save(data *models.TbCoinTask, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	}
	return dao.Insert(data)
}

func (dao *CoinTaskDao) Delete(id int) error {
	data := &models.TbCoinTask{Id: uint(id), SysStatus: 1}
	return dao.Save(data)
}
