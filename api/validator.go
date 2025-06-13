package api

import (
	"github.com/go-playground/validator/v10"
	"simplebank/util"
)

var validCurrency = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportCurrency(currency)
	}
	return false
}
