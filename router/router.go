package router

import (
	"blog-server/blog"
	"blog-server/ent"
	"github.com/gin-gonic/gin"
	"os"
)

func InitRouters(User *ent.User) *gin.Engine {
	env := os.Getenv("GIN_MODE")
	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)

	}
	blogservice := blog.NewBlogService()
	r := gin.New()
	v1 := r.Group("/api/v1")

	// POST a new blog
	v1.POST("/blog", gin.BasicAuth(gin.Accounts{User.Name: User.Password}), blogservice.AddBlog)
	// GET all blogs
	v1.GET("/blogs", blogservice.GetBlogs)
	// GET a blog by ID
	v1.GET("/blog/:id", blogservice.GetBlog)
	// PUT (update) a blog
	v1.PUT("/blog/:id", gin.BasicAuth(gin.Accounts{User.Name: User.Password}), blogservice.UpdateBlog)
	// DELETE a blog by ID
	v1.DELETE("/blog/:id", gin.BasicAuth(gin.Accounts{User.Name: User.Password}), blogservice.DeleteBlog)
	return r
}
