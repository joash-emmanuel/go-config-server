package main

import (
	"go-git/pull"
	"go-git/servicefetch"

	"github.com/gin-gonic/gin"
)

// func Loadenv() {
// 	err := godotenv.Load(".env")
// 	fmt.Println("pulled env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// }

func main() {
	// Loadenv()
	// clone.Clone_configs()

	go pull.Pull_configs()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/:filename/:environment", servicefetch.Retrieve_data)

	router.Run(":8080")
}

// http://localhost:8089/test-service/uat

//https://github.com/davidlukac-wisi/go-repo-sync/blob/65e3e305d7cf367813675a8d8afc4acd818b4feb/main.go#L252
//https://github.com/go-git/go-git/tree/v5.16.2/_examples
