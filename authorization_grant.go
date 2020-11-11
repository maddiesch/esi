package esi

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AuthorizationGrant struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type AuthorizationGrantRequest struct {
	Type         string `json:"grant_type" validate:"required,oneof=refresh_token authorization_code"`
	Code         string `json:"code,omitempty" validate:"rfe=Type:authorization_code)"`
	RefreshToken string `json:"refresh_token,omitempty" validate:"rfe=Type:refresh_token)"`
}

func NewAuthorizationGrant(ctx context.Context, app *App, in AuthorizationGrantRequest) (*AuthorizationGrant, error) {
	if err := Validate.Struct(app); err != nil {
		return nil, err
	}
	if err := Validate.Struct(in); err != nil {
		return nil, err
	}

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	uri := &url.URL{
		Scheme: "https",
		Host:   AuthenticationHost,
		Path:   "/v2/oauth/token",
	}

	req, err := http.NewRequest("POST", uri.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%s:%s", app.ClientID, app.SecretKey)),
		),
	)

	resp, err := HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Error: %s", resp.Status)
	}

	defer resp.Body.Close()

	grant := &AuthorizationGrant{}

	return grant, json.NewDecoder(resp.Body).Decode(grant)
}
