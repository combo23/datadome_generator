package services

import (
	"net/http"

	"github.com/combo23/datadome_generator/internal/constants"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (t *APIHandler) GetSolves(c *gin.Context) {
	t.Logger.Info("Handling request", zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
	solves, err := t.DB.GetSolves(c.Query("apikey"))
	if err != nil {
		t.Logger.Error(err.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
		c.JSON(500, gin.H{
			"error": ErrInternalError.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"solves": solves,
	})
	t.Logger.Info("returning solves", zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
}

func (t *APIHandler) BanUser(c *gin.Context) {
	t.AdminAuth(c)
	if c.IsAborted() {
		return
	}
	t.Logger.Info("Handling ban request", zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
	err := t.DB.BanUser(c.Query("apikey"))
	if err != nil {
		t.Logger.Error(err.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
		c.JSON(500, gin.H{
			"error": ErrInternalError.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
	t.Logger.Info("Banned user", zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
}

func (t *APIHandler) AddUser(c *gin.Context) {
	t.AdminAuth(c)
	if c.IsAborted() {
		return
	}
	t.Logger.Info("Adding new user", zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
	userpayload := new(RegisterPayload)
	if err := c.ShouldBindJSON(&userpayload); err != nil {
		t.Logger.Error(err.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidPayload.Error()})
		return
	}
	apikey, err := t.DB.AddUser(userpayload.Name, userpayload.Solves)
	if err != nil {
		t.Logger.Error(err.Error(), zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInternalError.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"apikey": apikey,
	})
	t.Logger.Info("Added new user", zap.String("path", c.FullPath()), zap.String("user", c.GetHeader(constants.APIKeyHeader)), zap.String("IP", c.ClientIP()))
}
