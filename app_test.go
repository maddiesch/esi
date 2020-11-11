package esi_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestApp(t *testing.T) {
	t.Skip()

	t.Run("App.LoginURL", func(t *testing.T) {
		t.Run("given a state and multiple scopes", func(t *testing.T) {
			location := app.LoginURL("foo", "publicData", "esi-characterstats.read.v1")

			spew.Dump(location)
		})
	})
}
