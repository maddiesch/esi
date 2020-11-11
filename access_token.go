package esi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
)

// VerifyAccessTokenSignature performs validation of the access token
func VerifyAccessTokenSignature(ctx context.Context, accessToken string) error {
	_, err := jwt.Parse(accessToken, func(unsafe *jwt.Token) (interface{}, error) {
		unsafeKid, present := unsafe.Header["kid"]
		if !present {
			return nil, errors.New("missing required kid value in token header")
		}

		kid, ok := unsafeKid.(string)
		if !ok {
			return nil, errors.New("invalid key header type in token header")
		}

		keys, err := FetchKeySet(ctx)
		if err != nil {
			return nil, err
		}

		if key, ok := keys[kid]; ok {
			return key, nil
		}

		return nil, errors.New("access token signed by unknown key")
	})

	return err
}

// Character contains information about a Character
type Character struct {
	CharacterID        int64
	CharacterName      string
	TokenType          string
	CharacterOwnerHash string
}

// FetchCharacter fetches the character information for the accessToken
func FetchCharacter(ctx context.Context, accessToken string) (*Character, error) {
	location := &url.URL{
		Scheme: "https",
		Host:   AuthenticationHost,
		Path:   "/oauth/verify",
	}

	req, err := http.NewRequest("GET", location.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("downstream http error: %s", resp.Status)
	}

	defer resp.Body.Close()

	character := &Character{}

	return character, json.NewDecoder(resp.Body).Decode(character)
}
