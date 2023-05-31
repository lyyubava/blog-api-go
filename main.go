package main

import (
	"fmt"
	"mini-blog-go/controllers"
	"mini-blog-go/controllers/auth"
	"mini-blog-go/controllers/users_interactions"
	"mini-blog-go/models"
	"os"

	"github.com/gin-gonic/gin"
)

func bootstrapRouter() *gin.Engine {
	r := gin.Default()

	models.ConnectDatabase()
	router_posts := r.Group("/post")
	router_posts.POST("/create", controllers.CreatePost)
	router_posts.PUT("/edit", controllers.EditPost)
	router_posts.DELETE("/delete", controllers.DeletePost)
	router_posts.GET("/", controllers.GetPost)
	router_posts.GET("/all", controllers.GetAllPosts)

	router_auth := r.Group("/auth")
	router_auth.POST("/register", auth.Register)
	router_auth.POST("/login", auth.Login)

	r.GET("/health", controllers.Health)

	r.POST("/comment", users_interactions.Comment)
	r.POST("/like", users_interactions.Like)
	return r

}

func main() {
	router := bootstrapRouter()
	router.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("API_PORT")))
}
