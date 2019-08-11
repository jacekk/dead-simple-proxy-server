package routing

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func configureCors(router *gin.Engine) *gin.Engine {
	config := cors.DefaultConfig()
	origins := os.Getenv("ALLOW_ORIGINS")

	if origins == "" {
		config.AllowAllOrigins = true
	} else {
		splitted := strings.Split(origins, ",")
		config.AllowOrigins = splitted
	}

	router.Use(cors.New(config))

	return router
}

// SetupRouter - initializes Gin router, configures CORS, along with all routes.
func SetupRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := configureCors(gin.Default())

	router.GET("/proxy/:slug", getProxyBySlug)
	ping := router.Group("/ping")
	{
		ping.GET("", GetPingRoute)
		ping.POST("", PostPingRoute)
	}

	return router
}

// InitRouter - initializes Gin router and runs it on given port.
func InitRouter(port string) error {
	router := SetupRouter()
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	return nil
}
