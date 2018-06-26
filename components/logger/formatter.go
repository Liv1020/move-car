package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Formatter Formatter
type Formatter struct {
	Prefix string
}

// Format Format
func (t *Formatter) Format(e *logrus.Entry) ([]byte, error) {
	f := "[%s] [%s] [%s] [%s]"
	a := []interface{}{t.Prefix, e.Time.Format("2006-01-02 15:04:05"), e.Level, e.Message}

	for key, value := range e.Data {
		f += " [%s=%s]"
		a = append(a, key)
		a = append(a, value)
	}

	v := fmt.Sprintf(f+"\n", a...)

	return []byte(v), nil
}
