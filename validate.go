package esi

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validate provides validation for values
var Validate = validator.New()

func init() {
	Validate.RegisterValidation(`rfe`, func(fl validator.FieldLevel) bool {
		param := strings.SplitN(fl.Param(), `:`, 2)

		paramName := param[0]
		paramValue := param[1]

		var targetValue reflect.Value
		if fl.Parent().Kind() == reflect.Ptr {
			targetValue = fl.Parent().Elem().FieldByName(paramName)
		} else {
			targetValue = fl.Parent().FieldByName(paramName)
		}

		toInt := func(in string) int64 {
			i, err := strconv.ParseInt(in, 10, 64)
			if err != nil {
				panic(err)
			}
			return i
		}

		isEqual := func(a reflect.Value, b string) bool {
			switch a.Kind() {
			case reflect.String:
				return a.String() == b

			case reflect.Slice, reflect.Map, reflect.Array:
				return int64(a.Len()) == toInt(b)

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return a.Int() == toInt(b)

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				i, err := strconv.ParseUint(b, 10, 64)
				if err != nil {
					panic(err)
				}
				return a.Uint() == i

			case reflect.Float32, reflect.Float64:
				f, err := strconv.ParseFloat(b, 64)
				if err != nil {
					panic(err)
				}

				return a.Float() == f

			default:
				return false

			}
		}

		if !isEqual(targetValue, paramValue) {
			return true
		}

		field := fl.Field()

		switch field.Kind() {
		case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
			return !field.IsNil()
		default:
			_, _, nullable := fl.ExtractType(field)
			if nullable && field.Interface() != nil {
				return true
			}
			return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
		}
	})
}
