package core

import (
	routers "go-rest-starter/src/api/v1/routers"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// health check variables
var _lastHealthCheck time.Time = time.Now().Add(time.Duration(-5) * time.Minute)
var _lastHealthResult = gin.H{}

// InitializeServer inits server using dig
func InitializeServer(r *gin.Engine) {
	//	e.InitEmailConfig()

	r.Use(gin.Recovery())

	r.MaxMultipartMemory = 4 << 20

	// r.Static("/s", "./public")

	r.GET("/health", health())

	v1 := r.Group("/api/v1")
	routers.AddAppStatusRoutes(v1)
	routers.AddAuthRoutes(v1)

	r.Use(invalidRoutes())

	port := ":" + strconv.Itoa(GetAppConfig().AppPort)
	r.Run(port)
}

func health() gin.HandlerFunc {
	return func(c *gin.Context) {

		if _lastHealthCheck.After(time.Now().Add(time.Duration(-5) * time.Minute)) {
			c.JSON(http.StatusOK, _lastHealthResult)
			return
		}

		healthy := false
		if isRedisHealthy() && isDbConnectionHealthy() {
			healthy = true
		}
		_lastHealthCheck = time.Now()
		_lastHealthResult = gin.H{
			"healthy": healthy,
			"app": gin.H{
				"port":        _config.AppPort,
				"environment": _config.AppEnv,
			},
			"services": gin.H{
				"database": isDbConnectionHealthy(),
				"redis":    isRedisHealthy(),
			},
		}
		c.JSON(http.StatusOK, _lastHealthResult)
	}
}

func invalidRoutes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"status": "failed", "message": "You've lost in this universe"})
	}
}
