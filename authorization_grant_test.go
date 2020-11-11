package esi_test

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/maddiesch/esi"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthorizationGrant(t *testing.T) {
	t.Skip()

	t.Run("NewAuthorizationGrant given a code", func(t *testing.T) {
		grant, err := esi.NewAuthorizationGrant(context.Background(), app, esi.AuthorizationGrantRequest{
			Type: `authorization_code`,
			Code: `9oXzqD6hu0yoVHdoyZQz1g`,
		})

		assert.NoError(t, err)

		spew.Dump(grant)
	})
}
