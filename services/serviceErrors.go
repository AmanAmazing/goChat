package services

import (
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid username or password")
var ErrDatabaseError = errors.New("database Error")
