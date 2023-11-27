package service

import (
	"context"
	"user_growth/dao"
	"user_growth/models"
)

//
// @Description Service层的封装
// @Author 代码小学生王木木
// @Date 2023/11/27 16:16
//

type CoinTaskService struct {
	ctx            context.Context
	daoCoinTaskDao *dao.CoinTaskDao
}

func NewCoinTaskService(ctx context.Context) *CoinTaskService {
	return &CoinTaskService{
		ctx:            ctx,
		daoCoinTaskDao: dao.NewCoinTaskDao(ctx),
	}
}

func (s *CoinTaskService) Get(id int) (*models.TbCoinTask, error) {
	return s.daoCoinTaskDao.Get(id)
}

func (s *CoinTaskService) FindByUid(uid, page, size int) ([]models.TbCoinTask, int64, error) {
	return s.daoCoinTaskDao.FindByUid(uid, page, size)
}

func (s *CoinTaskService) FindAllPager(uid, page, size int) ([]models.TbCoinTask, int64, error) {
	return s.daoCoinTaskDao.FindAllPager(uid, page, size)
}

func (s *CoinTaskService) Insert(data *models.TbCoinTask) error {
	return s.daoCoinTaskDao.Insert(data)
}

func (s *CoinTaskService) Update(data *models.TbCoinTask, musColumns ...string) error {
	return s.daoCoinTaskDao.Update(data, musColumns...)
}

func (s *CoinTaskService) Save(data *models.TbCoinTask, musColumns ...string) error {
	return s.daoCoinTaskDao.Save(data, musColumns...)
}
