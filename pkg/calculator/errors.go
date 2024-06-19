package calculator

import "errors"

// ErrInvalidRow is returned when a row in the CSV file is invalid
var ErrInvalidRow = errors.New("invalid row")

// ErrInvalidDate is returned when the date is invalid
var ErrInvalidDate = errors.New("invalid date")

// ErrNoCookies is returned when there are no cookies on the given date
var ErrNoCookies = errors.New("no cookies found for the given date")

// ErrInvalidData is returned when the data is invalid or empty
var ErrInvalidData = errors.New("invalid data")
