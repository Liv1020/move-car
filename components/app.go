package components

import (
	"fmt"

	"log"

	"os"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// DEV DEV
	DEV = "dev"
	// TEST test
	TEST = "test"
	// PROD prod
	PROD = "prod"
)

type app struct {
	Config *Config
	DB     *gorm.DB
	Logger *log.Logger
}

// Config Config
type Config struct {
	Env  string
	Port int
	DB   struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
		Charset  string
		MaxIdle  int
		MaxOpen  int
	}
}

// App App
var App = &app{}

func init() {
	App.Config = &Config{}
	if err := configor.Load(App.Config, "conf/app.yml"); err != nil {
		panic(err)
	}

	App.Logger = log.New(os.Stdout, "[App]", log.LstdFlags)

	cdb := App.Config.DB
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=3s&parseTime=true&loc=Local", cdb.User, cdb.Password, cdb.Host, cdb.Port, cdb.Database, cdb.Charset)

	var err error
	if App.DB, err = gorm.Open("mysql", args); err != nil {
		panic(err)
	}

	// 关闭tableName自动复数
	App.DB.SingularTable(true)
	// 默认不打印日志
	App.DB.LogMode(false)

	App.DB.DB().SetMaxIdleConns(cdb.MaxIdle)
	App.DB.DB().SetMaxOpenConns(cdb.MaxOpen)

	if App.Config.Env != PROD {
		App.DB.LogMode(true)
	}
}
