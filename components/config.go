package components

import "github.com/jinzhu/configor"

// Config Config
var Config config

// Config Config
type config struct {
	Port int
	DB   db
}

// DB 数据库连接
type db struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Charset      string
}

func init() {
	configor.Load(&Config, "config.yml")
}
