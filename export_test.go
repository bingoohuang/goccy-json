package json

import (
	"github.com/bingoohuang/json/internal/errors"
)

var (
	NewSyntaxError    = errors.ErrSyntax
	NewMarshalerError = errors.ErrMarshaler
)
