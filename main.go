package main

import (
	"embed"

	appcfg "github.com/umanit/toggl-redmine/internal/app"
	"github.com/umanit/toggl-redmine/internal/logger"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	appcfg.CreateAppDir()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "toggl track - Redmine",
		Width:  800,
		Height: 770,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Logger: logger.Create(),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
