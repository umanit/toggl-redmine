package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/umanit/toggl-redmine/internal/cfg"
	"github.com/umanit/toggl-redmine/internal/redmine"
	"github.com/umanit/toggl-redmine/internal/toggltrack"
)

type Redmine struct {
	cfg *cfg.ApiConfig
}

type redmineCurrentUser struct {
	User struct {
		Id int `json:"id"`
	} `json:"user"`
}

type timeEntry struct {
	IssueId  int     `json:"issue_id"`
	SpentOn  string  `json:"spent_on"`
	Hours    float64 `json:"hours"`
	Comments string  `json:"comments"`
}

type timeEntryWrapper struct {
	TimeEntry timeEntry `json:"time_entry"`
}

// Prepare permet d’altérer la requête HTTP qui va être envoyée.
func (a *Redmine) Prepare(req *http.Request) {
	req.Header.Add("X-Redmine-API-Key", a.cfg.Key)
}

// CheckUser vérifie si l’API est joignable et répond correctement.
func (a *Redmine) CheckUser(ctx context.Context) error {
	body, err := call(a, ctx, http.MethodGet, a.cfg.Url+"/users/current.json", nil)
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

// LoadTimeEntries va charger les temps déjà enregistrés entre deux dates fournies.
func (a *Redmine) LoadTimeEntries(ctx context.Context, dateFrom, dateTo time.Time) ([]redmine.TimeEntry, error) {
	body, err := call(a, ctx, http.MethodGet, a.cfg.Url+"/time_entries.json?user_id=me&limit=100&spent_on=><"+
		dateFrom.Format(time.DateOnly)+"|"+dateTo.Format(time.DateOnly), nil)
	if err != nil {
		return nil, err
	}

	var entries redmine.TimeEntriesList
	if err = json.Unmarshal(body, &entries); err != nil {
		return nil, err
	}

	return entries.TimeEntries, nil
}

// SynchronizeTimeEntries va créer de nouvelles entrées à partir de tâches toggl track.
func (a *Redmine) SynchronizeTimeEntries(ctx context.Context, tasks []toggltrack.AppTask) error {
	for _, t := range tasks {
		if !t.IsSyncable() {
			continue
		}

		te := &timeEntryWrapper{
			TimeEntry: timeEntry{
				IssueId:  t.Issue,
				SpentOn:  t.Date.Format(time.DateOnly),
				Hours:    t.DecimalDuration,
				Comments: t.Comment,
			},
		}

		jsonData, err := json.Marshal(te)
		if err != nil {
			return err
		}

		callCtx, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
		defer cancel()

		if _, err = call(a, callCtx, http.MethodPost, a.cfg.Url+"/time_entries.json", jsonData); err != nil {
			return err
		}
	}

	return nil
}

func NewRedmine(cfg *cfg.ApiConfig) *Redmine {
	return &Redmine{
		cfg: cfg,
	}
}
