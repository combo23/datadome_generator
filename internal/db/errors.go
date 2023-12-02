package db

import "fmt"

var (
	ErrInvalidAPIKey = fmt.Errorf("invalid api key")
	ErrNoSolves      = fmt.Errorf("no solves left")
	ErrInternalError = fmt.Errorf("internal error")
	ErrNoAPIKey      = fmt.Errorf("no api key provided")
)
