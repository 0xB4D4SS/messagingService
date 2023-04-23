package main

import "errors"

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty input")

// ErrNotFound is returned when no data found in db
var ErrNotFound = errors.New("Not found")
