package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Service interface {
	Prepare(req *http.Request)
	CheckUser(ctx context.Context) error
}

type CredentialsTest struct {
	TogglTrackOk bool
	RedmineOk    bool
}

func call(a Service, ctx context.Context, method, endpoint string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint, nil)
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
		return nil, fmt.Errorf("api call returned non-200 status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
