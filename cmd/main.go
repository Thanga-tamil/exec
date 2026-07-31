package main

import (
    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Thanga-tamil/exec/internal/config"
	"github.com/Thanga-tamil/exec/internal/router"
	logger "github.com/Thanga-tamil/exec/internal/utils"
	consts "github.com/Thanga-tamil/exec/internal/utils"
)


func main(){

	logger.Init("exec.log")
	
	addr := config.LoadConfig(consts.Path)

	serve := gin.Default()

    serve.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"https://cli-ent-react.thangatamil1177.workers.dev"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

	cal := serve.Group("/api/cal")

	router.Route(cal)

	serve.Run(addr)

}
