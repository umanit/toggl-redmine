package main

import (
	"context"
	"time"

	"github.com/umanit/toggl-redmine/internal/api"
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

// LoadConfig charge la configuration actuelle de l’application
func (a *App) LoadConfig() *cfg.Config {
	c, ok := cfg.ConfigFromContext(a.ctx)
	if !ok {
		return nil
	}
	return &c
}

// SaveConfig enregistre les informations fournies dans le formulaire de la page « Configurer »
func (a *App) SaveConfig(config cfg.Config) bool {
	c, ok := cfg.ConfigFromContext(a.ctx)
	if !ok {
		return false
	}

	if err := c.Save(config); err != nil {
		return false
	}

	return true
}

func (a *App) TestCredentials() api.CredentialsTest {
	cr := api.CredentialsTest{}
	c, ok := cfg.ConfigFromContext(a.ctx)
	if !ok {
		return cr
	}

	r := api.NewRedmine(c.Redmine)
	t := api.NewTogglTrack(c.Toggl)
	timeout := 3 * time.Second

	// On n’utilise pas le même contexte afin de ne pas arrêter automatiquement l’autre test si l’un des deux plante.
	rctx, rcancel := context.WithTimeout(context.Background(), timeout)
	defer rcancel()

	if err := r.CheckUser(rctx); err == nil {
		cr.RedmineOk = true
	}

	tctx, tcancel := context.WithTimeout(context.Background(), timeout)
	defer tcancel()

	if err := t.CheckUser(tctx); err == nil {
		cr.TogglTrackOk = true
	}

	return cr
}
