package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "example.com/go-docker-app/cryptit" // Ensure this matches your module name
)

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Define a "Hello, World!" endpoint
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": cryptit.EncryptIt("hey Please Encrypt it"), // Updated to use the correct package
        })
    })

    // Define another endpoint
    router.GET("/second-route", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello from the second route ok yes ooo ",
        })
    })

    // Define another API endpoint
    router.GET("/api/greet", func(c *gin.Context) {
        name := c.Query("name") // Get query parameter "name"
        if name == "" {
            name = "Guest"
        }
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, " + name + "!",
        })
    })

    // Run the server on port 8080
    router.Run(":8080")
}