package controllers

import (
	"meta_blog_api/initializers"
	"meta_blog_api/models"
	"meta_blog_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		Author      string
		Image       string
	}

	if c.Bind(&body) != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Failed to read body")
		return
	}

	if body.Title == "" || body.Description == "" || body.Image == "" || body.Author == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Title, Description and Author are required")
		return
	}

	user, _ := c.Get("user")

	newBlog := models.Blog{Title: body.Title, Image: body.Image, Author: body.Author, UserID: user.(models.User).ID}
	result := initializers.DB.Create(&newBlog).Preload("User")

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "error", "Failed to create blog")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Blog created successfully", newBlog)
}

func GetBlogs(c *gin.Context) {
	var blogs []models.Blog

	initializers.DB.Find(&blogs)

	utils.SuccessResponse(c, http.StatusOK, "Blogs fetched successfully", blogs)
}

func GetBlog(c *gin.Context) {
	var blog models.Blog

	if initializers.DB.First(&blog, c.Param("id")).Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Not Found", "Blog not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Blog fetched successfully", blog)
}

func UpdateBlog(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		Author      string
		Image       string
	}

	if c.Bind(&body) != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Failed to read body")
		return
	}

	var blog models.Blog

	if initializers.DB.First(&blog, c.Param("id")).Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Not Found", "Blog not found")
		return
	}

	initializers.DB.Model(&blog).Updates(models.Blog{Title: body.Title, Description: body.Description, Author: body.Author, Image: body.Image})
	utils.SuccessResponse(c, http.StatusOK, "Blog updated successfully", blog)
}

func DeleteBlog(c *gin.Context) {
	var blog models.Blog

	if initializers.DB.First(&blog, c.Param("id")).Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Not Found", "Blog not found")
		return
	}

	initializers.DB.Delete(&blog)
	utils.SuccessResponse(c, http.StatusOK, "Blog deleted successfully", nil)
}

func GetBlogsByAuthor(c *gin.Context) {
	var blogs []models.Blog

	initializers.DB.Where("author = ?", c.Param("author")).Find(&blogs)
	utils.SuccessResponse(c, http.StatusOK, "Blogs fetched successfully", blogs)
}

func GetUserBlogs(c *gin.Context) {
	var blogs []models.Blog
	user_id := c.Param("id")
	user, _ := c.Get("user")
	num64, _ := strconv.ParseUint(user_id, 10, 32)
	num := uint(num64)
	if num != user.(models.User).ID {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "You are not authorized to access this resource")
		return
	}

	initializers.DB.Where("user_id = ?", user.(models.User).ID).Preload("User").Find(&blogs)
	utils.SuccessResponse(c, http.StatusOK, "Blogs fetched successfully", blogs)
}

func SearchBlogs(c *gin.Context) {
	// Get the authenticated user from the context
	user, exists := c.Get("user")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "bad request", "User not found")
		return
	}

	// Extract user ID from the user object
	user_id := user.(models.User).ID

	// Query parameter for search filter
	search := c.Query("search")

	// Define a slice to hold the blogs
	var blogs []models.Blog

	// Construct the base query with user_id
	query := initializers.DB.Where("user_id = ?", user_id)

	// Add conditions for title and/or content filters if provided
	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Execute the query
	if err := query.Find(&blogs).Error; err != nil {
		// Handle database error
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch blogs", err.Error())
		return
	}

	// Check if no blogs were found
	if len(blogs) == 0 {
		// Respond with an empty array
		utils.SuccessResponse(c, http.StatusOK, "No blogs found", []models.Blog{})
		return
	}

	// Respond with the blogs found
	utils.SuccessResponse(c, http.StatusOK, "Blogs fetched successfully", blogs)
}
