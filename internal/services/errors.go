package services

import "fmt"

var (
	ErrInvalidAPIKey   = fmt.Errorf("invalid api key")
	ErrPayloadCreation = fmt.Errorf("error while creating payload")
	ErrInvalidPayload  = fmt.Errorf("invalid payload")
	ErrInvalidSite     = fmt.Errorf("invalid site")
	ErrInternalError   = fmt.Errorf("internal error")
)
