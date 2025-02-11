package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Define a "Hello, World!" endpoint
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World! eta joss na !?? again!!",
        })
    })
    router.GET("/second-route", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello from the second route ok yes ",
        })
    })

    // Define another endpoint
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