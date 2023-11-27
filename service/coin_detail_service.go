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

type CoinDetailService struct {
	ctx              context.Context
	daoCoinDetailDao *dao.CoinDetailDao
}

func NewCoinDetailService(ctx context.Context) *CoinDetailService {
	return &CoinDetailService{
		ctx:              ctx,
		daoCoinDetailDao: dao.NewCoinDetailDao(ctx),
	}
}

func (s *CoinDetailService) Get(id int) (*models.TbCoinDetail, error) {
	return s.daoCoinDetailDao.Get(id)
}

func (s *CoinDetailService) FindByUid(uid, page, size int) ([]models.TbCoinDetail, int64, error) {
	return s.daoCoinDetailDao.FindByUid(uid, page, size)
}

func (s *CoinDetailService) FindAllPager(uid, page, size int) ([]models.TbCoinDetail, int64, error) {
	return s.daoCoinDetailDao.FindAllPager(uid, page, size)
}

func (s *CoinDetailService) Insert(data *models.TbCoinDetail) error {
	return s.daoCoinDetailDao.Insert(data)
}

func (s *CoinDetailService) Update(data *models.TbCoinDetail, musColumns ...string) error {
	return s.daoCoinDetailDao.Update(data, musColumns...)
}

func (s *CoinDetailService) Save(data *models.TbCoinDetail, musColumns ...string) error {
	return s.daoCoinDetailDao.Save(data, musColumns...)
}
