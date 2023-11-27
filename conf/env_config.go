package conf

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
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
		Engine          string `json:"engine"`          // mysql
		Username        string `json:"username"`        // root
		Password        string `json:"password"`        // 123456
		Host            string `json:"host"`            // 127.0.0.1
		Port            int    `json:"port"`            // 3306
		Database        string `json:"database"`        // user_growth
		Charset         string `json:"charset"`         // utf8mb4
		ShowSQL         bool   `json:"showSQL"`         // false
		MaxIdleConns    int    `json:"maxIdleConns"`    // 2
		MaxOpenConns    int    `json:"maxOpenConns"`    // 10
		ConnMaxLifetime int    `json:"connMaxLifetime"` // 30m
	} `json:"db"`
}

func LoadConfigs(path string) {
	//LoadEnvConfig()
	//LoadJsonConfig()
	LoadYmlConfig(path)
}

func LoadEnvConfig() {
	pc := &ProjectConfig{}
	// load from os env
	if strConfigs := os.Getenv(envConfigName); len(strConfigs) > 0 {
		if err := json.Unmarshal([]byte(strConfigs), pc); err != nil {
			log.Fatalf("config.LoadEnvConfig(%s),error = %v\n", envConfigName, err)
			return
		}
	} else {
		print("参数不存在")
	}
	GlobalConfig = pc
}

func LoadJsonConfig() {
	pc := &ProjectConfig{}
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(configData, pc); err != nil {
		log.Fatalf("config.LoadJsonConfig(%s),error = %v\n", envConfigName, err)
	}
	GlobalConfig = pc
}

func LoadYmlConfig(path string) {
	pc := &ProjectConfig{}
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(pc); err != nil {
		panic(err)
	}
	GlobalConfig = pc
}
