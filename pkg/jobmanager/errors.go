package jobmanager

import (
	"errors"
)

var (
	JobExistsError    = errors.New("job exists")
	JobNotFoundError  = errors.New("job not found")
	InvalidJobIDError = errors.New("invalid job id")
	Unauthenticated   = errors.New("Unauthenticated")
	InvalidArgument   = errors.New("invalid argument")
)
