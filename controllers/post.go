package controllers

import (
	"mini-blog-go/models"
	"mini-blog-go/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type EditPostInput struct {
	Id    uint   `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type DeletePostInput struct {
	Id uint `json:"id" binding:"required"`
}

type GetPostInput struct {
	Id uint `json:"id" binding:"required"`
}

func CreatePost(c *gin.Context) {
	t, err := token.ExtractTokenId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	models.DB.Model(models.User{}).Where("id = ?", t).Take(&user)
	createdBy := user.Username
	post := models.Post{Title: input.Title, Body: input.Body, UserID: t, CreatedBy: createdBy}
	models.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func EditPost(c *gin.Context) {
	u_id, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input EditPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//post := models.Post{Title: input.Title, Body: input.Body, UserID: t}
	var post models.Post
	models.DB.Model(models.Post{}).Where("id = ?", input.Id).Take(&post)
	if post.UserID != u_id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	models.DB.Model(&models.Post{}).Where("id = ?", input.Id).Update("title", input.Title)
	models.DB.Model(&models.Post{}).Where("id = ?", input.Id).Update("body", input.Body)
	c.JSON(http.StatusCreated, gin.H{"data": "success"})
}

func DeletePost(c *gin.Context) {
	u_id, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input DeletePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var post models.Post
	models.DB.Model(models.Post{}).Where("id = ?", input.Id).Take(&post)
	if post.UserID != u_id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	models.DB.Model(&models.Post{}).Where("id = ?", input.Id).Delete(input.Id)
	c.JSON(http.StatusCreated, gin.H{"data": "success"})
}

func GetPost(c *gin.Context) {
	var input GetPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var post models.Post
	models.DB.Model(models.Post{}).Where("id = ?", input.Id).Take(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})

}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}
