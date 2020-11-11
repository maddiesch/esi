package esi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-jose/go-jose"
)

// KeySet contains a collection of JWT keys used to validate ESI access tokens
type KeySet map[string]interface{}

// FetchKeySet fetches all the keys from the authorization server
func FetchKeySet(ctx context.Context) (KeySet, error) {
	location := &url.URL{
		Scheme: "https",
		Host:   AuthenticationHost,
		Path:   "/oauth/jwks",
	}

	req, err := http.NewRequest("GET", location.String(), nil)
	if err != nil {
		return KeySet{}, err
	}

	resp, err := HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return KeySet{}, err
	}
	defer resp.Body.Close()

	body := struct {
		Keys []jose.JSONWebKey `json:"keys"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return KeySet{}, err
	}

	set := KeySet{}
	for _, key := range body.Keys {
		set[key.KeyID] = key.Key
	}

	return set, nil
}
