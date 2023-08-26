package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grealyve/blog-post-api/controllers"
)

func BlogRoute(router *gin.Engine) {
	blogGroup := router.Group("/blog")

	blogGroup.GET("/post/:post_id", controllers.GetByID)
	blogGroup.GET("/post", controllers.GetAllPosts)
	blogGroup.POST("/post", controllers.NewPost)
	blogGroup.PUT("/post/:post_id", controllers.UpdatePost)
	blogGroup.DELETE("/post/:post_id", controllers.DeletePost)
}
