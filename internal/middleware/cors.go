package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"https://cli-ent-react.thangatamil1177.workers.dev"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}
