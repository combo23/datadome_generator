package services

import (
	"net/http"

	"github.com/combo23/datadome_generator/internal/constants"

	"github.com/combo23/datadome_generator/internal/datadome"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (t *APIHandler) Datadome(c *gin.Context) {
	t.Logger.Info("Handling request", zap.String("path", c.FullPath()), zap.String("user", c.GetHeader("X-api-key")), zap.String("IP", c.ClientIP()))
	t.Auth(c)
	if c.IsAborted() {
		return
	}
	userpayload := new(DataDomePayload)
	if err := c.ShouldBindJSON(&userpayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidPayload.Error()})
		return
	}
	for _, x := range []string{userpayload.Site, userpayload.UserAgent, userpayload.Cid, userpayload.ResponsePage, userpayload.IP, userpayload.Referer} {
		if x == "" {
			t.Logger.Error(ErrInvalidPayload.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidPayload.Error()})
			return
		}
	}
	if !validateSite(userpayload.Site) {
		t.Logger.Error(ErrInvalidSite.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidSite.Error()})
		return
	}

	task := datadome.DataDomePayload{
		Host:         userpayload.Site,
		UserAgent:    userpayload.UserAgent,
		Cid:          userpayload.Cid,
		ResponsePage: userpayload.ResponsePage,
		IP:           userpayload.IP,
		Referer:      userpayload.Referer,
	}
	task.JsType = "ch"
	payloadch, err := task.GetPayload()
	if err != nil {
		t.Logger.Error(err.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
		c.JSON(500, gin.H{
			"error": ErrPayloadCreation.Error(),
		})
		return
	}

	task.JsType = "le"
	payloadle, err := task.GetPayload()
	if err != nil {
		t.Logger.Error(err.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
		c.JSON(500, gin.H{
			"error": ErrPayloadCreation.Error(),
		})
		return
	}
	err = t.DB.UpdateUsage(c.GetHeader(constants.APIKeyHeader))
	if err != nil {
		t.Logger.Warn(err.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": ErrInternalError.Error(),
		})
		return //not fatal
	}

	c.JSON(200, gin.H{
		"ch": payloadch,
		"le": payloadle,
	})
	t.Logger.Info("return dd payload", zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
}
