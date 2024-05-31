package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/umanit/toggl-redmine/internal/cfg"
)

type TogglTrack struct {
	cfg *cfg.ApiConfig
}

type togglTrackCurrentUser struct {
	Id int `json:"id"`
}

func (a *TogglTrack) Authenticate(req *http.Request) {
	req.SetBasicAuth(a.cfg.Key, "api_token")
}

func (a *TogglTrack) CheckUser(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.cfg.Url+"/me", nil)
	if err != nil {
		return err
	}

	a.Authenticate(req)

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

func NewTogglTrack(cfg *cfg.ApiConfig) *TogglTrack {
	return &TogglTrack{
		cfg: cfg,
	}
}
