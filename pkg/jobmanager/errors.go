package jobmanager

import (
	"errors"
)

var (
	ErrJobExists       = errors.New("job exists")
	ErrJobNotFound     = errors.New("job not found")
	ErrInvalidJobID    = errors.New("invalid job id")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrInvalidArgument = errors.New("invalid argument")
)
