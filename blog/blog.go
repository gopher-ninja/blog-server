package blog

import (
	"blog-server/ent"
	"blog-server/ent/blog"
	"blog-server/ent/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

type blogService struct {
	dbClient *ent.Client
}

type BlogService interface {
	GetBlogs(*gin.Context)
	GetBlog(*gin.Context)
	AddBlog(*gin.Context)
	DeleteBlog(*gin.Context)
	UpdateBlog(*gin.Context)
}

func NewBlogService() BlogService {
	return &blogService{}
}

func dbinit() *ent.Client {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUserName := os.Getenv("DB_USER")
	client, err := ent.Open("postgres", "host="+dbHost+" port="+dbPort+" user="+dbUserName+" dbname="+dbName+" password="+dbPassword+" sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to bloggres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func CreateAdminUser(ctx context.Context) (*ent.User, error) {
	client := dbinit()
	defer client.Close()
	adminName := os.Getenv("ADMIN_NAME")
	if adminName == "" {
		adminName = "admin"
	}
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "Admin@123"
	}
	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminPassword == "" {
		adminPassword = "admin@gmail.com"
	}
	existAdmin, err := client.User.Query().Where(user.Name(adminName)).Only(context.Background())
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, fmt.Errorf("failed creating user: %w", err)
		}
	} else if existAdmin.Name == adminName {
		log.Println("Admin user already exist: ", existAdmin.String())
		return existAdmin, nil
	}
	id := uuid.New()
	u, err := client.User.
		Create().
		SetUserId(id).
		SetName(adminName).
		SetEmail(adminEmail).
		SetPassword(adminPassword).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("Admin user created successfully: ", u.String())
	return u, nil
}

// Add blog
func (bs *blogService) AddBlog(c *gin.Context) {
	client := dbinit()
	defer client.Close()
	var blog ent.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	id := uuid.New()
	blogDetail, err := client.Blog.Create().SetBlogId(id).
		SetBlogTitle(blog.BlogTitle).
		SetBlogContent(blog.BlogContent).Save(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blogDetail": blogDetail})
}

// Get all blogs
func (bs *blogService) GetBlogs(c *gin.Context) {
	client := dbinit()
	defer client.Close()
	blogs, err := client.Blog.Query().All(context.Background())
	if err != nil {
		log.Printf("error getting user: %s", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

// Get single blog using ID
func (bs *blogService) GetBlog(c *gin.Context) {
	client := dbinit()
	defer client.Close()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blogDetail, err := client.Blog.Query().Where(blog.BlogId(id)).Only(context.Background())
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blogDetail": blogDetail})
}

// Update Blog
func (bs *blogService) UpdateBlog(c *gin.Context) {
	client := dbinit()
	defer client.Close()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameter"})
	}
	var blogDetail ent.Blog
	err = c.BindJSON(&blogDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}
	err = client.Blog.Update().SetBlogTitle(blogDetail.BlogTitle).SetBlogContent(blogDetail.BlogContent).Where(blog.BlogId(id)).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed while updating blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blogDetail": blogDetail})
}

// Delete Blog
func (bs *blogService) DeleteBlog(c *gin.Context) {
	client := dbinit()
	defer client.Close()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BlogId"})
		return
	}
	_, err = client.Blog.Delete().Where(blog.BlogId(id)).Exec(context.Background())
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Blog deleted successfully"})
}
