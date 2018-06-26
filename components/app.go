package components

import (
	"fmt"

	"github.com/Liv1020/move-car/components/logger"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

type app struct {
	db     *gorm.DB
	config *Config
	logger *logrus.Logger
}

// App App
var App = &app{}

// DB DB
func (t *app) DB() *gorm.DB {
	return t.db.New()
}

// Config Config
func (t *app) Config() *Config {
	return t.config
}

// Logger Logger
func (t *app) Logger() *logrus.Logger {
	return t.logger
}

func init() {
	App.config = &Config{}
	if err := configor.Load(App.config, "conf/app.yml"); err != nil {
		panic(err)
	}
	gin.SetMode(App.config.Mode)

	App.logger = logrus.New()
	App.logger.Formatter = &logger.Formatter{
		Prefix: "GIN-app",
	}

	cdb := App.config.DB
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=3s&parseTime=true&loc=Local", cdb.User, cdb.Password, cdb.Host, cdb.Port, cdb.Database, cdb.Charset)
	App.logger.Infof("数据库连接：%s", args)

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

	if App.config.Mode != gin.ReleaseMode {
		App.db.LogMode(true)
		App.db.SetLogger(&logger.DbLogger{})
	}
}
