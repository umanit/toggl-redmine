package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/umanit/toggl-redmine/internal/cfg"
)

type Redmine struct {
	cfg *cfg.ApiConfig
}

type redmineCurrentUser struct {
	User struct {
		Id int `json:"id"`
	} `json:"user"`
}

func (a *Redmine) Authenticate(req *http.Request) {
	req.Header.Add("X-Redmine-API-Key", a.cfg.Key)
}

func (a *Redmine) CheckUser(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.cfg.Url+"/users/current.json", nil)
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
		return fmt.Errorf("redmine returned non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
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

func NewRedmine(cfg *cfg.ApiConfig) *Redmine {
	return &Redmine{
		cfg: cfg,
	}
}
