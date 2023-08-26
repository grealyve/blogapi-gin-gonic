package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/grealyve/blog-post-api/models"
)

func GetByID(c *gin.Context) {
	postId := c.Param("post_id")
	num, _ := strconv.ParseInt(postId, 10, 64)

	p, _ := models.GetByID(num)
	c.JSON(http.StatusOK, p)
}

func GetAllPosts(c *gin.Context) {
	posts := models.GetAllPosts()
	c.Bind(posts)

	c.JSON(http.StatusOK, posts)
}

func NewPost(c *gin.Context) {
	var p models.Post
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.CreatePost()

	c.JSON(http.StatusOK, p)
}

func DeletePost(c *gin.Context) {
	postId := c.Param("post_id")
	num, _ := strconv.ParseInt(postId, 10, 64)

	p := models.DeletePost(num)

	c.JSON(http.StatusOK, p)
}

func UpdatePost(c *gin.Context) {
	postId := c.Param("post_id")
	var updatePost = &models.Post{}
	num, _ := strconv.ParseInt(postId, 10, 64)
	c.Bind(updatePost)

	post, db := models.GetByID(num)
	if updatePost.Title != "" {
		post.Title = updatePost.Title
	}

	if updatePost.Content != "" {
		post.Content = updatePost.Content
	}

	if updatePost.Username != "" {
		post.Username = updatePost.Username
	}
	db.Save(post)

	c.JSON(http.StatusOK, post)
}
