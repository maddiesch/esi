package esi

import (
	"net/url"
	"strings"
)

// App is an ESI application. Can be created & managed at https://developers.eveonline.com
type App struct {
	ClientID    string   `validate:"required"`
	SecretKey   string   `validate:"required"`
	CallbackURL *url.URL `validate:"required"`
}

// LoginURL returns a login URL with the specified scopes
func (a *App) LoginURL(state string, scopes ...string) *url.URL {
	query := url.Values{
		"client_id":     {a.ClientID},
		"redirect_uri":  {a.CallbackURL.String()},
		"response_type": {"code"},
		"scope":         {strings.Join(scopes, " ")},
		"state":         {state},
	}

	return &url.URL{
		Scheme:   "https",
		Host:     AuthenticationHost,
		Path:     "/v2/oauth/authorize",
		RawQuery: query.Encode(),
	}
}
