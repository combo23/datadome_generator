package routes

import (
	"net/http"

	"github.com/combo23/datadome_generator/internal/services"
	"github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {
	handler, err := services.SetupDependencies()
	if err != nil {
		panic(err)
	}
	handler.Logger.Info("Successfully setup dependencies")
	engine.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "alive",
		})
	})
	engine.POST("/datadome", handler.Datadome)
	engine.POST("/register", handler.AddUser)
	engine.GET("/solves", handler.GetSolves)
	engine.POST("/ban", handler.BanUser)
}
