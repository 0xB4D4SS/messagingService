package errors

import "errors"

// ErrEmpty is returned when input data is empty
var ErrEmpty = errors.New(`{"response": "error", "errors": "Empty input"}`)

// ErrNotFound is returned when no data found in db
var ErrNotFound = errors.New(`{"response": "error", "errors": "Not found"}`)

// ErrCouldNotInsert is returned when could not insert data into db
var ErrCouldNotInsert = errors.New(`{"response": "error", "errors": "Could not insert"}`)
