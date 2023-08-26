package models

import (
	"github.com/grealyve/blog-post-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

type Post struct {
	ID       uint   `gorm:"primary_key"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func init() {
	db = config.ConnectDB()
	db.AutoMigrate(&Post{})
}

func (p *Post) CreatePost() *Post {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllPosts() []Post {
	var posts []Post
	db.Find(&posts)
	return posts
}

func GetByID(Id int64) (*Post, *gorm.DB) {
	var post Post
	db.Where("ID=?", Id).Find(&post)
	return &post, db
}

func DeletePost(Id int64) Post {
	var post Post
	db.Where("ID=?", Id).Delete(post)
	return post
}
