package api

import (
	"fmt"
	"os"

	"github.com/combo23/datadome_generator/internal/routes"
	"github.com/gin-gonic/gin"
)

func Serve(releasemode string) {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.Setup(r)
	switch releasemode {
	case "dev":
		fmt.Println("Starting server in dev mode")
		gin.SetMode(gin.DebugMode)
		r.Run(os.Getenv("ADDRESS"))
	case "production":
		fmt.Println("Starting server in production mode")
		gin.SetMode(gin.ReleaseMode)
		r.Run(os.Getenv("ADDRESS"))
	default:
		panic("invalid release mode: " + releasemode)
	}
}
