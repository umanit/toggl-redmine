package api

import (
	"bytes"
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

// call effectue un appel HTTP à une des APIs et renvoie le body résultant ou une erreur en cas de problème.
func call(a Service, ctx context.Context, method, endpoint string, payload []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	a.Prepare(req)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if payload != nil {
		req.Body = io.NopCloser(bytes.NewReader(payload))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf("api call returned non-200/201 status code: %d; body %s; payload %s",
			resp.StatusCode,
			string(body),
			string(payload),
		)
	}

	return io.ReadAll(resp.Body)
}
