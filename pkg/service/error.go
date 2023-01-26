package service

import (
	"errors"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrAlreadyExists      = errors.New("already exists")
	ErrAccountDoesntExist = errors.New("accountId doesnt exist")
)
