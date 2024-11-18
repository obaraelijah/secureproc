package jobmanager

import (
	"errors"
)

var (
	JobExistsError    = errors.New("job exists")
	JobNotFoundError  = errors.New("job not found")
	InvalidJobIDError = errors.New("invalid job id")
)
