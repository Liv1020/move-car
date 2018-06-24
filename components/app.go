package components

import (
	"fmt"

	"log"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type app struct {
	db     *gorm.DB
	Config *Config
	Logger *log.Logger
}

// Config Config
type Config struct {
	Mode string
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
	Wechat struct {
		AppID     string
		AppSecret string
		OAuthUrl  string
	}
	FrontedUrl string
	Aliyun     struct {
		AccessKeyId     string
		AccessKeySecret string
	}
}

// App App
var App = &app{}

// DB 防止污染db
func (t *app) DB() *gorm.DB {
	return t.db.New()
}

func init() {
	App.Config = &Config{}
	if err := configor.Load(App.Config, "conf/app.yml"); err != nil {
		panic(err)
	}
	gin.SetMode(App.Config.Mode)

	App.Logger = log.New(os.Stdout, "[App]", log.LstdFlags)

	cdb := App.Config.DB
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=3s&parseTime=true&loc=Local", cdb.User, cdb.Password, cdb.Host, cdb.Port, cdb.Database, cdb.Charset)

	var err error
	if App.db, err = gorm.Open("mysql", args); err != nil {
		panic(err)
	}

	// 关闭tableName自动复数
	App.db.SingularTable(true)
	// 默认不打印日志
	App.db.LogMode(false)
	// 跳过关联保存
	App.db.Set("gorm:save_associations", false)

	App.db.DB().SetMaxIdleConns(cdb.MaxIdle)
	App.db.DB().SetMaxOpenConns(cdb.MaxOpen)

	if App.Config.Mode != gin.ReleaseMode {
		App.db.LogMode(true)
	}
}
