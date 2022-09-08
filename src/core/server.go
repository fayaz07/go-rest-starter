package core

import (
	routers "go-rest-starter/src/api/v1/routers"
	config "go-rest-starter/src/core/config"
	log "go-rest-starter/src/core/logger"
	"go-rest-starter/src/utils/constants"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// health check variables
var _lastHealthCheck time.Time = time.Now().Add(time.Duration(-5) * time.Minute)
var _lastHealthResult = gin.H{}

// InitializeServer inits server using dig
func InitializeServer(r *gin.Engine) {
	//	e.InitEmailConfig()

	log.I("Starting GIN Server...")
	setGinMode()

	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1", "http://127.0.0.1:7001"},
		AllowMethods:     []string{"PUT", "POST"},
		AllowHeaders:     []string{"Origin", "Host"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.MaxMultipartMemory = 4 << 20

	// r.Static("/s", "./public")

	r.GET("/health", health())

	v1 := r.Group("/api/v1")
	routers.AddAppStatusRoutes(v1)
	routers.AddAuthRoutes(v1)

	r.Use(invalidRoutes())

	port := ":" + strconv.Itoa(config.GetAppConfig().AppPort)
	r.Run("localhost" + port)
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
				"port":        config.GetAppConfig().AppPort,
				"environment": config.GetAppConfig().AppEnv,
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

func setGinMode() {
	switch config.GetAppConfig().AppEnv {
	case constants.PROD_ENV:
		gin.SetMode(gin.ReleaseMode)
	case constants.DEV_ENV:
		gin.SetMode(gin.DebugMode)
	case constants.TEST_ENV, constants.STAGING_ENV:
		gin.SetMode(gin.TestMode)
	}
}
