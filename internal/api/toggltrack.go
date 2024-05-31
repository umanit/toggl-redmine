package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.cfg.Url+"/me", nil)
	if err != nil {
		return err
	}

	a.Prepare(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("toggl track returned non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
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

func (a *TogglTrack) LoadTasks(ctx context.Context, dateFrom, dateTo time.Time) ([]toggltrack.ApiTask, error) {
	dateTo = dateTo.AddDate(0, 0, 1)

	dl := time.DateOnly
	urlParams := url.Values{
		"start_date": {dateFrom.Format(dl)},
		"end_date":   {dateTo.Format(dl)},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.cfg.Url+"/me/time_entries?"+urlParams.Encode(), nil)
	if err != nil {
		return nil, err
	}

	a.Prepare(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("toggl track returned non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tasks []toggltrack.ApiTask
	if err = json.Unmarshal(body, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func NewTogglTrack(cfg *cfg.ApiConfig) *TogglTrack {
	return &TogglTrack{
		cfg: cfg,
	}
}
