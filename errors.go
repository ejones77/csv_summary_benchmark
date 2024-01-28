package main

import "errors"

var (
	ErrInvalidColumn    = errors.New("column not in headers")
	ErrNoFile           = errors.New("file not found")
	ErrInvalidOperation = errors.New("invalid operation")
	ErrFileEmpty        = errors.New("file has no data")
	ErrInvalidFormat    = errors.New("invalid csv format")
)
