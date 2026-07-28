package jwthandler

import "errors"

var (
	ErrMissingJwtSecret = errors.New("missing jwt secret")
)
