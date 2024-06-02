package logger

import (
	"path/filepath"

	"github.com/umanit/toggl-redmine/internal/app"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

func Create() logger.Logger {
	appDir := app.GetAppDir()
	logFile := filepath.Join(appDir, "logs.log")
	return logger.NewFileLogger(logFile)
}
