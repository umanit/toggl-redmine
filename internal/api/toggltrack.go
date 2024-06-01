package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/umanit/toggl-redmine/internal/cfg"
	"github.com/umanit/toggl-redmine/internal/toggltrack"
)

type TogglTrack struct {
	cfg *cfg.ApiConfig
}

type togglTrackCurrentUser struct {
	Id int `json:"id"`
}

func (a *TogglTrack) Prepare(req *http.Request) {
	req.SetBasicAuth(a.cfg.Key, "api_token")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
}

func (a *TogglTrack) CheckUser(ctx context.Context) error {
	body, err := call(a, ctx, http.MethodGet, a.cfg.Url+"/me")
	if err != nil {
		return err
	}

	var cu togglTrackCurrentUser
	err = json.Unmarshal(body, &cu)
	if err != nil {
		return err
	}

	return nil
}

// LoadTasks va charger les tâches entre deux dates fournies.
func (a *TogglTrack) LoadTasks(ctx context.Context, dateFrom, dateTo time.Time) ([]toggltrack.ApiTask, error) {
	dateTo = dateTo.AddDate(0, 0, 1)
	dl := time.DateOnly
	urlParams := url.Values{
		"start_date": {dateFrom.Format(dl)},
		"end_date":   {dateTo.Format(dl)},
	}

	body, err := call(a, ctx, http.MethodGet, a.cfg.Url+"/me/time_entries?"+urlParams.Encode())
	if err != nil {
		return nil, err
	}

	var tasks []toggltrack.ApiTask
	if err = json.Unmarshal(body, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// HasRunningTask vérifie s’il y a une tâche en cours sur toggl track afin de ne pas la synchroniser par inadvertence.
func (a *TogglTrack) HasRunningTask(ctx context.Context) bool {
	body, err := call(a, ctx, http.MethodGet, a.cfg.Url+"/me/time_entries/current")
	if err != nil {
		return false
	}

	// L’API de toggl track renvoie « null » s’il n’y a pas de tâche en cours au lieu d’une vraie réponse JSON
	return string(body) != "null"
}

func NewTogglTrack(cfg *cfg.ApiConfig) *TogglTrack {
	return &TogglTrack{
		cfg: cfg,
	}
}
