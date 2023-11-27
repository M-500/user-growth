package db_helper

// @Description
// @Author 代码小学生王木木
// @Date 2023/11/27 14:01
import (
	"fmt"
	"log"
	"time"
	"user_growth/conf"
	"xorm.io/xorm"
)

var dbEngine *xorm.Engine

func InitDb() {
	if dbEngine != nil {
		return
	}
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		conf.GlobalConfig.Db.Username,
		conf.GlobalConfig.Db.Password,
		conf.GlobalConfig.Db.Host,
		conf.GlobalConfig.Db.Port,
		conf.GlobalConfig.Db.Database,
		conf.GlobalConfig.Db.Charset,
	)
	if engine, err := xorm.NewEngine(conf.GlobalConfig.Db.Engine, sourceName); err != nil {
		log.Fatalf("db_helper.InitDb(%s),%v\n", sourceName, err)
	} else {
		dbEngine = engine
	}

	if conf.GlobalConfig.Db.MaxIdleConns > 0 {
		dbEngine.SetMaxIdleConns(conf.GlobalConfig.Db.MaxIdleConns)
	}

	if conf.GlobalConfig.Db.MaxOpenConns > 0 {
		dbEngine.SetMaxOpenConns(conf.GlobalConfig.Db.MaxOpenConns)
	}
	if conf.GlobalConfig.Db.ConnMaxLifetime > 0 {
		dbEngine.SetConnMaxLifetime(time.Minute * time.Duration(conf.GlobalConfig.Db.ConnMaxLifetime))
	}
}

func GetDb() *xorm.Engine {
	return dbEngine
}
