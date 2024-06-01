package main

import (
	"context"
	"fmt"
	"time"

	"github.com/umanit/toggl-redmine/internal/api"
	"github.com/umanit/toggl-redmine/internal/cfg"
	"github.com/umanit/toggl-redmine/internal/toggltrack"
)

const (
	httpCallTimeout = 3 * time.Second
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

	// On n’utilise pas le même contexte afin de ne pas arrêter automatiquement l’autre test si l’un des deux plante.
	rctx, rcancel := context.WithTimeout(context.Background(), httpCallTimeout)
	defer rcancel()

	if err := r.CheckUser(rctx); err == nil {
		cr.RedmineOk = true
	}

	tctx, tcancel := context.WithTimeout(context.Background(), httpCallTimeout)
	defer tcancel()

	if err := t.CheckUser(tctx); err == nil {
		cr.TogglTrackOk = true
	}

	return cr
}

// LoadTasks permet de charger les temps enregistrés sur toggl track et de les afficher dans un tableau afin de décider
// lesquels seront à synchroniser sur Redmine.
func (a *App) LoadTasks(dateFromStr, dateToStr string) *toggltrack.AskedTasks {
	c, ok := cfg.ConfigFromContext(a.ctx)
	if !ok {
		return nil
	}

	dl := time.RFC3339

	dateFrom, err := time.Parse(dl, dateFromStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	dateTo, err := time.Parse(dl, dateToStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	t := api.NewTogglTrack(c.Toggl)
	ctx, cancel := context.WithTimeout(context.Background(), httpCallTimeout)
	defer cancel()

	tasks, err := t.LoadTasks(ctx, dateFrom, dateTo)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &toggltrack.AskedTasks{
		Entries:        toggltrack.ProcessTasks(tasks),
		HasRunningTask: t.HasRunningTask(ctx),
	}
}
