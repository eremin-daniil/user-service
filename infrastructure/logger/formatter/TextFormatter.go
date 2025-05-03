package formatter

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
	"user_service/infrastructure/logger"
)

type TextFormatter struct {
}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timeFormat := entry.Time.Format(time.RFC822)
	level := strings.ToUpper(entry.Level.String())
	requestID := entry.Data[logger.RequestIDFieldKey]
	if requestID != nil {
		return []byte(fmt.Sprintf("%s [%s] (%s): %s\n", timeFormat, level, requestID, entry.Message)), nil
	}
	return []byte(fmt.Sprintf("%s [%s]: %s\n", timeFormat, level, entry.Message)), nil
}
