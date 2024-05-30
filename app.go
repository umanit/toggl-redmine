package main

import (
	"context"

	"github.com/umanit/toggl-redmine/internal/cfg"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = cfg.ContextWithConfig(ctx)
}

// CanSynchronize indique si l’accès à l’écran « Synchroniser » est possible ou non.
// Il est nécessaire d’avoir configuré les clés et URLs des APIs pour y avoir accès.
func (a *App) CanSynchronize() bool {
	c, ok := cfg.ConfigFromContext(a.ctx)
	if !ok {
		return false
	}

	return c.AllFill()
}
