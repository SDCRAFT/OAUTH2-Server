package main

import (
	"github.com/gin-gonic/gin"
	_ "sdcraft.fun/oauth2/database"
	_ "sdcraft.fun/oauth2/globals"
	Routes "sdcraft.fun/oauth2/routes"
)

func main() {
	router := gin.Default()
	Routes.Register_v1_routes(router.Group("/api/v1"))
	router.Run(":8080")
}
