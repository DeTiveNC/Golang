package controllers

import (
	"github.com/detivenc/jsoncrudapi/initializers/db"
	"github.com/detivenc/jsoncrudapi/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := db.DB.Create(&post) // pass pointer of data to

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	db.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the id
	id := c.Param("id")
	// Get the post you need
	var post models.Post

	db.DB.First(&post, id)

	// Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id of the post
	id := c.Param("id")
	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find the post
	var post models.Post

	db.DB.First(&post, id)

	// Update it
	db.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id
	id := c.Param("id")
	// Delete
	db.DB.Delete(&models.Post{}, id)
	// Return
	c.JSON(200, gin.H{
		"Delete": "Successfull",
	})
}
