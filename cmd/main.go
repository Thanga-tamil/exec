package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Thanga-tamil/exec/internal/config"
	"github.com/Thanga-tamil/exec/internal/router"
	"github.com/Thanga-tamil/exec/internal/middleware"
	logger "github.com/Thanga-tamil/exec/internal/utils"
	consts "github.com/Thanga-tamil/exec/internal/utils"
)


func main(){

	logger.Init("exec.log")
	
	addr := config.LoadConfig(consts.Path)

	serve := gin.Default()

	serve.Use(middleware.CORS())

	cal := serve.Group("/api/cal")

	router.Route(cal)

	serve.Run(addr)

}
