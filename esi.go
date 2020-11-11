// Package esi provides methods for interacting with the EVE Online ESI API
//
// Read More: https://developers.eveonline.com
// API Docs: https://esi.evetech.net/ui/
package esi

import (
	"net/http"
	"time"
)

// AuthenticationHost is the hostname for the authorization server
var AuthenticationHost = `login.eveonline.com`

// HTTPClient is the HTTP client that will be used to perform all requests to the authorization server
var HTTPClient = &http.Client{
	Timeout: 5 * time.Second,
}
