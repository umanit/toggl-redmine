package api

import (
	"context"
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
