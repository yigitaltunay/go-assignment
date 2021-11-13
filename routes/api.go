package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yigitaltunay/go-assignment/app/handlers"
	"github.com/yigitaltunay/go-assignment/config"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func InitializeRoutes() {
	
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/currency/:type", handlers.Currency)
	r.GET("/test", handlers.Test)
	port := os.Getenv("PORT")
	if port == "" {
		r.Run(":" + config.BackendConfig.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
	r.Run(":" + port)

}
