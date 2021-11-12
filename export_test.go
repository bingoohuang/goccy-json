package json

import (
	"github.com/bingoohuang/go-json/internal/errors"
)

var (
	NewSyntaxError    = errors.ErrSyntax
	NewMarshalerError = errors.ErrMarshaler
)
