package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tpSpace/go-jwt/controllers"
	"github.com/tpSpace/go-jwt/initializers"
	"github.com/tpSpace/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth ,controllers.Validate)

	r.Run()
}