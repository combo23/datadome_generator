package services

import (
	"os"

	"github.com/combo23/datadome_generator/internal/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (t *APIHandler) Auth(c *gin.Context) {
	apikey := c.GetHeader(constants.APIKeyHeader)
	isvalid, err := t.DB.ValidateAPIKey(apikey)
	if err != nil {
		t.Logger.Error("Error validating API Key", zap.Error(err))
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	if !isvalid {
		t.Logger.Warn("Invalid API Key", zap.String("IP", apikey))
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.Next()
}

func (t *APIHandler) AdminAuth(c *gin.Context) {
	if c.GetHeader(constants.APIKeyHeader) != os.Getenv("ADMIN_API_KEY") {
		t.Logger.Warn("Invalid API Key", zap.String("IP", c.GetHeader(constants.APIKeyHeader)))
		c.JSON(403, gin.H{
			"error": ErrInvalidAPIKey.Error(),
		})
		c.Abort()
		return
	}
	c.Next()
}
