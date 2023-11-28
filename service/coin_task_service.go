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

func (s *CoinTaskService) GetByTask(task string) (*models.TbCoinTask, error) {
	return s.daoCoinTaskDao.GetByTask(task)
}

func (s *CoinTaskService) FindAll() ([]models.TbCoinTask, error) {
	return s.daoCoinTaskDao.FindAll()
}

func (s *CoinTaskService) Save(data *models.TbCoinTask, musColumns ...string) error {
	return s.daoCoinTaskDao.Save(data, musColumns...)
}
