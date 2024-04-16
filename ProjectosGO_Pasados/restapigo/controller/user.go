package controller

import (
	"github.com/detivenc/restapigo/initializers"
	"github.com/detivenc/restapigo/model"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	if initializers.DB.Find(&users).Error != nil {
		c.JSON(400, gin.H{
			"users": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"users": users,
	})
}

func PostUser(c *gin.Context) {
	// Get data for the req body
	var body struct {
		Name     string
		Email    string
		Password string
	}

	c.Bind(&body)
	// Create post
	userCreate := model.User{Name: body.Name, Email: body.Email, Password: body.Password}
	result := initializers.DB.Create(&userCreate)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"user": userCreate,
	})
}

func DeleteUser(c *gin.Context) {
	// Obtain the param to eliminate
	id := c.Param("id")
	// Eliminate user
	result := initializers.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Delete Unsuccessful",
		})
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"user": "Delete Succesful",
	})
}
func UpdateUser(c *gin.Context) {
	// Get id
	id := c.Param("id")
	// Get data for the req body
	var body struct {
		Name     string
		Email    string
		Password string
	}

	c.Bind(&body)
	// Find the user
	var userUpdate model.User

	initializers.DB.Find(&userUpdate, id)

	// Update User
	initializers.DB.Model(&userUpdate).Updates(model.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})
	// Return it
	c.JSON(200, gin.H{
		"user": userUpdate,
	})
}
