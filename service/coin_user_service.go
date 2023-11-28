package service

import (
	"context"
	"user_growth/dao"
	"user_growth/models"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/28 9:52
//

type CoinUserService struct {
	ctx         context.Context
	daoCoinUser *dao.CoinUserDao
}

func NewCoinUserService(ctx context.Context) *CoinUserService {
	return &CoinUserService{
		ctx:         ctx,
		daoCoinUser: dao.NewCoinUserDao(ctx),
	}
}

func (s *CoinUserService) Get(id int) (*models.TbCoinUser, error) {
	return s.daoCoinUser.Get(id)
}

// GetByUid get model by uid
func (s *CoinUserService) GetByUid(uid int) (*models.TbCoinUser, error) {
	return s.daoCoinUser.GetByUid(uid)
}

// FindAllPager get all models
func (s *CoinUserService) FindAllPager(page, size int) ([]models.TbCoinUser, int64, error) {
	return s.daoCoinUser.FindAllPager(page, size)
}

// Save with Insert and Update
func (s *CoinUserService) Save(data *models.TbCoinUser, musColumns ...string) error {
	return s.daoCoinUser.Save(data, musColumns...)
}
