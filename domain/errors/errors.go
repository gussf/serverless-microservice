package errors

import "errors"

var (
	ErrNoCarsAvailable       = errors.New("there are no available cars")
	ErrFetchingAvailableCars = errors.New("an error occurred when fetching available cars")
)
