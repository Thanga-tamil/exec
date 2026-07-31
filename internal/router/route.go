package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Thanga-tamil/exec/internal/handler"
)

func Route(group *gin.RouterGroup){

	group.GET("/calculate", handler.Calculate)

}

