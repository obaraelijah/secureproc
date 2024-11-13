package jobmanager

import (
	"errors"
	"fmt"
)

type JobExistsError struct {
	jobName string
}

func NewJobExistsError(jobName string) error {
	return &JobExistsError{jobName: jobName}
}

func (j *JobExistsError) Error() string {
	return fmt.Sprintf("job with name '%s' already exists", j.jobName)
}

func (j *JobExistsError) Is(err error) bool {
	_, ok := err.(*JobExistsError)

	return err != nil && ok
}

type JobNotFoundError struct {
	jobID string
}

func NewJobNotFoundError(jobID string) error {
	return &JobNotFoundError{jobID: jobID}
}

func (j *JobNotFoundError) Error() string {
	return fmt.Sprintf("job with ID '%s' not found", j.jobID)
}

func (j *JobNotFoundError) Is(err error) bool {
	_, ok := err.(*JobNotFoundError)

	errors.Is(nil, nil)

	return err != nil && ok
}

type InvalidJobID struct {
	jobID string
}

func NewInvalidJobID(jobID string) error {
	return &InvalidJobID{jobID: jobID}
}

func (j *InvalidJobID) Error() string {
	return fmt.Sprintf("'%s' is not a valid job ID", j.jobID)
}

func (j *InvalidJobID) Is(err error) bool {
	_, ok := err.(*InvalidJobID)

	errors.Is(nil, nil)

	return err != nil && ok
}
