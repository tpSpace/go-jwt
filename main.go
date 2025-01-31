package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tpSpace/go-jwt/controllers"
	"github.com/tpSpace/go-jwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)

	r.Run()
}