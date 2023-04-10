package auth

import (
	"mini-blog-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}
	user.Username = input.Username
	user.Password = input.Password
	user.FullName = input.FullName
	user.Email = input.Email
	user_existance := models.DB.Model(models.User{}).Where("username = ?", user.Username).Take(&user)

	if user_existance.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must me unique within the system"})
		return
	}

	_, err := user.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User has been successfully registered"})

}
