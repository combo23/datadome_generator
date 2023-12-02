package services

import (
	"os"
	"strings"

	"github.com/combo23/datadome_generator/internal/constants"
	"github.com/combo23/datadome_generator/internal/db"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupDependencies() (APIHandler, error) {
	handler := new(APIHandler)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"logs/API.log", "stdout"},
		ErrorOutputPaths: []string{"logs/API-Errors.log", "stderr"},
		EncoderConfig:    encoderConfig,
		Development:      true, //change in prod
	}

	logger, err := cfg.Build()
	if err != nil {
		return *handler, err
	}
	handler.Logger = logger
	handler.Logger.Sync()
	handler.DB = db.CreateDBInstance(os.Getenv("MONGO_URI"))
	handler.DB.Connect()
	return *handler, nil
}

func validateSite(site string) bool {
	for x := range constants.SupportedSites {
		if strings.Contains(site, x) {
			return true
		}
	}
	return false
}
