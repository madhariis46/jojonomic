package main

import (
	r "document-service/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()
	r.Router(router)
	router.Run(":" + port)
}

func init() {
	godotenv.Load(".env")
	//c.RedisInit()
}
