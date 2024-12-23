package controllers

import (
	"meta_blog_api/initializers"
	"meta_blog_api/models"
	"meta_blog_api/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Register(c *gin.Context) {
	//Get the request body
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	if body.Name == "" || body.Email == "" || body.Password == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Name, Email and Password are required")
		return
	}

	//HashPassword
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	//check if the email already exists
	var user models.User
	userExist := initializers.DB.Where("email = ?", body.Email).First(&user)
	if userExist.Error == nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Email already exists")
		return
	}

	//Add the new user to the database
	newUser := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", result.Error.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User created successfully", newUser)
}

func Login(c *gin.Context) {
	//Get the email and Password from the request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Failed to read body")
		return
	}

	//Email and Password is required
	if body.Email == "" || body.Password == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Email and Password are required")
		return
	}

	//get the user from the database
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	//Check if the user exists
	if user.ID == 0 {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid email or password")
		return
	}

	//validate the user password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid email or password")
		return
	}

	//generate for the user a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	user.Token = tokenString
	initializers.DB.Save(&user)

	userResponse := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Token:     user.Token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "User logged in successfully", userResponse)
}

func ValidateAuth(c *gin.Context) {
	//Get the user from the context
	user, _ := c.Get("user")
	utils.SuccessResponse(c, http.StatusOK, "User is authenticated", user)
}

func ResetPassword(c *gin.Context) {
	//Get the data from the request body
	var body struct {
		Email       string
		NewPassword string
	}

	if c.Bind(&body) != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Failed to read body")
		return
	}

	if body.Email == "" || body.NewPassword == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "Email and NewPassword are required")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		utils.ErrorResponse(c, http.StatusBadRequest, "bad request", "User not found")
		return
	}
	user.Password = string(hash)
	initializers.DB.Save(&user)
	utils.SuccessResponse(c, http.StatusOK, "Password reset successfully", nil)
}
