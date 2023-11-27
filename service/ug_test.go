package service

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
	"user_growth/common"
	"user_growth/conf"
	"user_growth/db_helper"
	"user_growth/models"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/27 16:30
//

func initDb() {
	time.Local = time.UTC
	conf.LoadConfigs("../etc/dev.yaml")
	db_helper.InitDb()
}

func TestCoinTaskService_Save(t *testing.T) {
	initDb()
	s := NewCoinTaskService(context.Background())
	data := &models.TbCoinTask{
		Id:         0,
		Task:       "publish article",
		Coin:       10,
		Limit:      3,
		Start:      common.Now(),
		SysUpdated: common.Now(),
		SysCreated: common.Now(),
	}

	if err := s.Save(data); err != nil {
		t.Errorf("Save(%+v) error= %v", data, err)
	} else {
		log.Printf("Save data=%+v\n", data)
	}
}
