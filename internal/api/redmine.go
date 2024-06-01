package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/umanit/toggl-redmine/internal/cfg"
	"github.com/umanit/toggl-redmine/internal/redmine"
)

type Redmine struct {
	cfg *cfg.ApiConfig
}

type redmineCurrentUser struct {
	User struct {
		Id int `json:"id"`
	} `json:"user"`
}

func (a *Redmine) Prepare(req *http.Request) {
	req.Header.Add("X-Redmine-API-Key", a.cfg.Key)
}

func (a *Redmine) CheckUser(ctx context.Context) error {
	body, err := call(a, ctx, http.MethodGet, a.cfg.Url+"/users/current.json")
	if err != nil {
		return err
	}

	var cu redmineCurrentUser
	err = json.Unmarshal(body, &cu)
	if err != nil {
		return err
	}

	return nil
}

func (a *Redmine) LoadTimeEntries(ctx context.Context, dateFrom, dateTo time.Time) ([]redmine.TimeEntry, error) {
	body, err := call(a, ctx, http.MethodGet, a.cfg.Url+"/time_entries.json?user_id=me&limit=100&spent_on=><"+
		dateFrom.Format(time.DateOnly)+"|"+dateTo.Format(time.DateOnly))
	if err != nil {
		return nil, err
	}

	var entries redmine.TimeEntriesList
	if err = json.Unmarshal(body, &entries); err != nil {
		return nil, err
	}

	return entries.TimeEntries, nil
}

func NewRedmine(cfg *cfg.ApiConfig) *Redmine {
	return &Redmine{
		cfg: cfg,
	}
}
