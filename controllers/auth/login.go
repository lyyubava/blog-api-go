package auth

import (
	"fmt"
	"mini-blog-go/models"
	"mini-blog-go/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json"username" binding:"required"`
	Password string `json"password" binding:"required"`
}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}

	user.Username = input.Username
	user.Password = input.Password

	token, err := models.LoginVerify(user.Username, user.Password)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
