package users_interactions

import (
	"fmt"
	"mini-blog-go/models"
	"mini-blog-go/utils/notify"
	"mini-blog-go/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeInput struct {
	PostID uint `json:"postId" binding:"required"`
}

func Like(c *gin.Context) {
	var input LikeInput
	user_id, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	like := models.Like{}
	like.PostID = input.PostID
	like.CreatedBy = user_id

	var count int64
	models.DB.Where("post_id = ? AND created_by = ?", input.PostID, user_id).Find(&models.Like{}).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You already liked a post"})

	}
	query_res := models.DB.Create(&like).Error
	if query_res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The post you want to like does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "successfully like post"})

	post := models.Post{}
	_ = models.DB.First(&post, input.PostID)
	user := models.User{Username: post.CreatedBy}
	models.DB.First(&user)

	username := user.Username
	email := user.Email
	likeEvent := notify.Event{}
	likeEvent.EventInfo.EventName = "Like"
	likeEvent.EventInfo.EventTime = like.CreatedAt
	likeEvent.EventInfo.EventUser = username
	likeEvent.EventInfo.EventUserEmail = email

	likeFormatted := fmt.Sprintf("%s liked your post", username)
	likeEvent.EventDetails = likeFormatted

	notify.Publish(likeEvent)

}
