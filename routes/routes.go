package routes

import (
	"MyTransactAPP/controllers"
	"MyTransactAPP/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Application URLs
	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", controllers.Logout)
	}

	// User routes
	user := r.Group("/user").Use(middleware.AuthMiddleware())
	{
		user.PUT("/update", controllers.UpdateUser)
	}

	// Payment routes
	r.GET("/payments/confirm/:id", controllers.ConfirmPayment)
	payments := r.Group("/payments").Use(middleware.AuthMiddleware())
	{
		payments.POST("/", controllers.CreatePayment)
		payments.GET("/:id", controllers.GetTransactionDetails)
	}

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
