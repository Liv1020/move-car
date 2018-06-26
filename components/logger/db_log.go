package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// DbLogger DbLogger
type DbLogger struct {
}

var dbLog *logrus.Logger

// Print Print
func (t *DbLogger) Print(v ...interface{}) {
	f := ""
	for i := 0; i < len(v); i++ {
		f += "%v "
	}
	dbLog.Print(fmt.Sprintf(f, v...))
}

func init() {
	dbLog = logrus.New()
	dbLog.Formatter = &DbFormatter{
		Prefix: "GIN-gorm",
	}
}
