package esi_test

import (
	"testing"

	"github.com/maddiesch/esi"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Skip()

	t.Run("rfe", func(t *testing.T) {
		type bReq struct {
			A string
			B string `validate:"rfe=A:foo"`
		}

		t.Run("given a matching target value", func(t *testing.T) {
			err := esi.Validate.Struct(bReq{
				A: "foo",
				B: "",
			})

			assert.Error(t, err)
		})

		t.Run("given a non-matching target value", func(t *testing.T) {
			err := esi.Validate.Struct(bReq{
				A: "bar",
				B: "",
			})

			assert.NoError(t, err)
		})
	})
}
