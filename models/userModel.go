package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Blogs    []Blog
}

type Blog struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UserID      uint   // Foreign key referencing the ID field of the User struct
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Define the relationship
}
