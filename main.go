package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// Use gib to generate a random string
		randomString := "23"
		c.String(200, fmt.Sprintf("Hello, World i want to check if the update sys is working! Here's a random string: %s", randomString))
	})

	fmt.Println("Starting server on on on :8080...")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
