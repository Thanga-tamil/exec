package main

import (
	"github.com/Thanga-tamil/exec/internal/router"
	logger "github.com/Thanga-tamil/exec/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	Addr = "0.0.0.0:6969"
)
func main(){

	logger.Init("exec.log")
	
	serve := gin.Default()

	cal := serve.Group("/api/cal")

	router.Route(cal)

	serve.Run(Addr)

}
