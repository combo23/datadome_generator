package services

import (
	"github.com/combo23/datadome_generator/internal/db"
	"go.uber.org/zap"
)

type APIHandler struct {
	Logger *zap.Logger
	DB     *db.Database
}

type DataDomePayload struct {
	Site         string `json:"site"`
	UserAgent    string `json:"user-agent"`
	Cid          string `json:"cid"`
	ResponsePage string `json:"responsePage"`
	IP           string `json:"ip"`
	Referer      string `json:"referer"`
}

type RegisterPayload struct {
	Name   string `json:"name"`
	Solves int    `json:"solves"`
}
