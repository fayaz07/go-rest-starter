package core

import (
	routers "go-rest-starter/src/api/v1/routers"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InitializeServer inits server using dig
func InitializeServer(r *gin.Engine) {
	//	e.InitEmailConfig()

	r.Use(gin.Recovery())

	r.MaxMultipartMemory = 4 << 20

	// r.Static("/s", "./public")

	v1 := r.Group("/v1")
	routers.AddAppStatusRoutes(v1)

	r.Use(invalidRoutes())

	port := ":" + strconv.Itoa(GetAppConfig().AppPort)
	r.Run(port)
}

func invalidRoutes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"status": "failed", "message": "You've lost in this universe"})
	}
}
