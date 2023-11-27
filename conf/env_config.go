package conf

import (
	"encoding/json"
	"log"
	"os"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/27 13:53
//

var GlobalConfig *ProjectConfig

const envConfigName = "USE_GROWTH_CONFIG"

type ProjectConfig struct {
	Db struct {
		Engine          string // mysql
		Username        string // root
		Password        string // 123456
		Host            string // 127.0.0.1
		Port            int    // 3306
		Database        string // user_growth
		Charset         string // utf8mb4
		ShowSQL         bool   // false
		MaxIdleConns    int    // 2
		MaxOpenConns    int    // 10
		ConnMaxLifetime int    // 30m
	}
}

func LoadConfigs() {
	LoadEnvConfig()
}

func LoadEnvConfig() {
	pc := &ProjectConfig{}
	// load from os env
	if strConfigs := os.Getenv(envConfigName); len(strConfigs) > 0 {
		if err := json.Unmarshal([]byte(strConfigs), pc); err != nil {
			log.Fatalf("config.LoadEnvConfig(%s),error = %v\n", envConfigName, err)
			return
		}
	}
	GlobalConfig = pc
}
