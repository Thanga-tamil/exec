package main

import (
	"github.com/Thanga-tamil/exec/internal/config"
	"github.com/Thanga-tamil/exec/internal/router"
	logger "github.com/Thanga-tamil/exec/internal/utils"
	consts "github.com/Thanga-tamil/exec/internal/utils"
	"github.com/gin-gonic/gin"
)


func main(){

	logger.Init("exec.log")
	
	addr := config.LoadConfig(consts.Path)

	serve := gin.Default()

	cal := serve.Group("/api/cal")

	router.Route(cal)

	serve.Run(addr)

}
