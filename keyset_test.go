package esi_test

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/maddiesch/esi"
	"github.com/stretchr/testify/require"
)

func TestFetchKeySet(t *testing.T) {
	set, err := esi.FetchKeySet(context.Background())

	require.NoError(t, err)

	spew.Dump(set)
}
