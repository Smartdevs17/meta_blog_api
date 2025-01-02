package main

import (
	"fmt"
	"meta_blog_api/controllers"
	"meta_blog_api/initializers"
	"meta_blog_api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("Server started running successfully")
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Authentication routes
	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)
	r.GET("/api/auth/validate", middleware.AuthMiddleware, controllers.ValidateAuth)
	r.POST("/api/auth/resetpassword", controllers.ResetPassword)

	// Blog routes
	r.POST("/api/blogs", middleware.AuthMiddleware, controllers.CreateBlog)
	r.GET("/api/blogs", controllers.GetBlogs)
	r.GET("/api/blogs/search", controllers.SearchBlogs)
	r.GET("/api/blogs/single/:id", controllers.GetBlog)
	r.GET("/api/blogs/user/:id", middleware.AuthMiddleware, controllers.GetUserBlogs)
	r.PUT("/api/blogs/:id", middleware.AuthMiddleware, controllers.UpdateBlog)
	r.DELETE("/api/blogs/:id", middleware.AuthMiddleware, controllers.DeleteBlog)

	// User routes
	r.GET("/api/users", middleware.AuthMiddleware, controllers.GetAllUsers)

	r.Run() // Listen and serve on localhost:3000
}
