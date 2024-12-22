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
	fmt.Println("Server stated running successfully")
	r := gin.Default()
	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)
	r.GET("/api/auth/validate", middleware.AuthMiddleware, controllers.ValidateAuth)
	r.POST("/api/auth/resetpassword", controllers.ResetPassword)
	r.POST("/api/blogs", middleware.AuthMiddleware, controllers.CreateBlog)
	r.GET("/api/blogs", controllers.GetBlogs)
	r.GET("/api/blogs/search", controllers.SearchBlogs)
	r.GET("/api/blogs/single/:id", controllers.GetBlog)
	r.GET("/api/blogs/user/:id", middleware.AuthMiddleware, controllers.GetUserBlogs)
	r.PUT("/api/blogs/:id", middleware.AuthMiddleware, controllers.UpdateBlog)
	r.DELETE("/api/blogs/:id", middleware.AuthMiddleware, controllers.DeleteBlog)

	r.Run() // listen and serve on localhost:3000
}
