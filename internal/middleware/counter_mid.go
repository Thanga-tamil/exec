package middleware

import (
	log "github.com/Thanga-tamil/exec/internal/utils"
	"github.com/gin-gonic/gin"
)

func Counter(counter *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		*counter++
		log.Info("request count: ", *counter)
		c.Next()
	}
}
