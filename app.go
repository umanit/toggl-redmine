package main

import (
	"context"
	"time"

	"github.com/umanit/toggl-redmine/internal/api"
	"github.com/umanit/toggl-redmine/internal/cfg"
	"github.com/umanit/toggl-redmine/internal/toggltrack"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	httpCallTimeout = 3 * time.Second
	errorEvent      = "goError"
)

type App struct {
	ctx context.Context
}

type LoadedConfig struct {
	Config  cfg.Config
	IsValid bool
}

// NewApp crée une nouvelle application de struct App.
func NewApp() *App {
	return &App{}
}

// startup est appelé quand l’application démarre. Le contexte est enregistré afin que l’on puisse appeler les méthodes
// du runtime.
func (a *App) startup(ctx context.Context) {
	a.ctx = cfg.ContextWithConfig(ctx)
}

func (a *App) logError(msg string) {
	runtime.EventsEmit(a.ctx, errorEvent)
	runtime.LogError(a.ctx, msg)
}

func (a *App) logErrorf(msg string, args ...interface{}) {
	runtime.EventsEmit(a.ctx, errorEvent)
	runtime.LogErrorf(a.ctx, msg, args)
}

func (a *App) logFatal(msg string) {
	runtime.EventsEmit(a.ctx, errorEvent)
	runtime.LogFatal(a.ctx, msg)
}

func (a *App) logFatalf(msg string, args ...interface{}) {
	runtime.EventsEmit(a.ctx, errorEvent)
	runtime.LogFatalf(a.ctx, msg, args)
}

// CanSynchronize indique si l’accès à l’écran « Synchroniser » est possible ou non.
// Il est nécessaire d’avoir configuré les clés et URLs des APIs pour y avoir accès.
func (a *App) CanSynchronize() bool {
	c := cfg.ConfigFromContext(a.ctx)

	return c.AllValuesFilled()
}

// LoadConfig charge la configuration actuelle de l’application en indiquant si elle est complète.
func (a *App) LoadConfig() LoadedConfig {
	c := cfg.ConfigFromContext(a.ctx)
	return LoadedConfig{
		Config:  c,
		IsValid: c.AllValuesFilled(),
	}
}

// SaveConfig enregistre les informations fournies dans le formulaire de la page « Configurer »
func (a *App) SaveConfig(config cfg.Config) {
	c := cfg.ConfigFromContext(a.ctx)

	if err := c.Save(config); err != nil {
		a.logFatal("can't save config")
	}
}

// TestCredentials vérifie que les APIs sont joignables.
func (a *App) TestCredentials() api.CredentialsTest {
	c := cfg.ConfigFromContext(a.ctx)
	r := api.NewRedmine(c.Redmine)
	t := api.NewTogglTrack(c.Toggl)
	cr := api.CredentialsTest{}

	// On n’utilise pas le même contexte afin de ne pas arrêter automatiquement l’autre test si l’un des deux échoue.
	rctx, rcancel := context.WithTimeout(context.Background(), httpCallTimeout)
	defer rcancel()

	err := r.CheckUser(rctx)
	if err != nil {
		a.logErrorf("can't check Redmine credentials %v", err)
	} else {
		cr.RedmineOk = true
	}

	tctx, tcancel := context.WithTimeout(context.Background(), httpCallTimeout)
	defer tcancel()

	err = t.CheckUser(tctx)
	if err != nil {
		a.logErrorf("can't check toggl track credentials %v", err)
	} else {
		cr.TogglTrackOk = true
	}

	return cr
}

// LoadTasks permet de charger les temps enregistrés sur toggl track et de les afficher dans un tableau afin de décider
// lesquels seront à synchroniser sur Redmine.
func (a *App) LoadTasks(dateFromStr, dateToStr string) *toggltrack.AskedTasks {
	c := cfg.ConfigFromContext(a.ctx)
	dl := time.RFC3339

	dateFrom, err := time.Parse(dl, dateFromStr)
	if err != nil {
		a.logErrorf("cannot parse dateFrom %s", dateFromStr)
		return nil
	}
	dateTo, err := time.Parse(dl, dateToStr)
	if err != nil {
		a.logErrorf("cannot parse dateTo %s", dateToStr)
		return nil
	}

	t := api.NewTogglTrack(c.Toggl)
	ctx, cancel := context.WithTimeout(context.Background(), httpCallTimeout)
	defer cancel()

	tasks, err := t.LoadTasks(ctx, dateFrom, dateTo)
	if err != nil {
		a.logErrorf("cannot load toggl track tasks %v", err)
		return nil
	}

	r := api.NewRedmine(c.Redmine)
	timeEntries, err := r.LoadTimeEntries(ctx, dateFrom, dateTo)
	if err != nil {
		a.logErrorf("cannot load Redmine time entries %v", err)
		return nil
	}

	h, err := t.HasRunningTask(ctx)
	if err != nil {
		a.logErrorf("cannot check running task on toggl track %v", err)
		return nil
	}

	return &toggltrack.AskedTasks{
		Entries:        toggltrack.ProcessTasks(tasks, timeEntries),
		HasRunningTask: h,
	}
}

func (a *App) SynchronizeTasks(tasks []toggltrack.AppTask) {
	if len(tasks) == 0 {
		return
	}

	c := cfg.ConfigFromContext(a.ctx)
	r := api.NewRedmine(c.Redmine)
	if err := r.SynchronizeTimeEntries(a.ctx, tasks); err != nil {
		a.logErrorf("cannot synchronize time entries to Redmine %v", err)
		return
	}

	return
}
