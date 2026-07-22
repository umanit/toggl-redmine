package logger

import (
	"github.com/umanit/toggl-redmine/internal/app"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

func Create() logger.Logger {
	return logger.NewFileLogger(app.GetLogsPath())
}
