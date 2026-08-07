package main

import (
	"strings"
	"github.com/gin-gonic/gin"

	"github.com/Thanga-tamil/exec/internal/config"
	"github.com/Thanga-tamil/exec/internal/router"
	"github.com/Thanga-tamil/exec/internal/middleware"
	logger "github.com/Thanga-tamil/exec/internal/utils"
	consts "github.com/Thanga-tamil/exec/internal/utils"
)


func main(){

	logger.Init("exec.log")
	
	addr, mode := config.LoadConfig(consts.Path)

	setMode(mode)

	serve := gin.Default()

	serve.Use(middleware.CORS())

	cal := serve.Group("/api/cal")

	router.Route(cal)

	serve.Run(addr)

}

func setMode(mode string){
	switch strings.ToLower(mode) {
	case "release", "production", "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev", "qa", "":
		gin.SetMode(gin.DebugMode)
	default: 
		gin.SetMode(gin.TestMode)
	}
}
