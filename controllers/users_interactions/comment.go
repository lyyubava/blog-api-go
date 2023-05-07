package users_interactions

import (
	"encoding/json"
	"mini-blog-go/models"
	"mini-blog-go/utils/notify"
	"mini-blog-go/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentInput struct {
	Comment string `json:"comment" binding:"required"`
	PostID  uint   `json:"postId" binding:"required"`
}

type CommentEventDetails struct {
	CommentId string
}

func Comment(c *gin.Context) {
	var input CommentInput
	user_id, err := token.ExtractTokenId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := models.Comment{}
	comment.Comment = input.Comment
	comment.PostID = input.PostID
	comment.CreatedBy = user_id
	query_res := models.DB.Create(&comment).Error
	if query_res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The post you want to comment on does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "comment was successfully added"})

	commentEventDetails := CommentEventDetails{CommentId: strconv.FormatUint(uint64(comment.ID), 10)}
	commentEventDetailsStr, _ := json.Marshal(&commentEventDetails)
	commentEvent := notify.Event{}
	commentEvent.EventInfo.EventName = "Comment"
	commentEvent.EventInfo.EventTime = comment.CreatedAt
	commentEvent.EventDetails = string(commentEventDetailsStr)

	notify.Publish(commentEvent)
	return

}
