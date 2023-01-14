package main

import (
	"blog-server/blog"
	"blog-server/router"
	"context"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Create admin user
	adminUser, err := blog.CreateAdminUser(context.Background())
	if err != nil {
		log.Fatalf("failed while creating admin user: %v", err)
	}
	// Init Router
	newRouter := router.InitRouters(adminUser)
	// Start Server
	err = newRouter.Run()
	if err != nil {
		log.Fatal("failed to run server")
	}
}
