package core

import (
	routers "go-rest-starter/src/api/v1/routers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InitializeServer inits server using dig
func InitializeServer(r *gin.Engine) {
	//	e.InitEmailConfig()

	r.Use(gin.Recovery())

	r.MaxMultipartMemory = 4 << 20

	// r.Static("/s", "./public")

	r.GET("/health", health())

	v1 := r.Group("/v1")
	routers.AddAppStatusRoutes(v1)

	r.Use(invalidRoutes())

	port := ":" + strconv.Itoa(GetAppConfig().AppPort)
	r.Run(port)
}

func health() gin.HandlerFunc {
	return func(c *gin.Context) {
		if didConnectToDB() {
			c.JSON(http.StatusOK,
				gin.H{
					"healthy": true,
					"app": gin.H{
						"port":        _config.AppPort,
						"environment": _config.AppEnv,
					},
					"services": gin.H{
						"database": didConnectToDB(),
					},
				},
			)
		} else {
			c.JSON(http.StatusOK, gin.H{"healthy": false, "services": gin.H{"database": didConnectToDB()}})
		}
	}
}

func invalidRoutes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"status": "failed", "message": "You've lost in this universe"})
	}
}
