package main

import (
	"go-git/clone"
	"go-git/microservicesfetch"
	"go-git/pull"

	"github.com/gin-gonic/gin"
)

func main() {
	clone.Clone_configs()
	pull.Pull_configs()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/:filename/:environment", microservicesfetch.Retrieve_data)

	router.Run(":8089")
}

// http://localhost:8089/test-service/uat
