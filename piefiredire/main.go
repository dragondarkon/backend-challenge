package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API with Golang and Gin"})
}

func main() {
	router := gin.Default()
	router.GET("/", HomeHandler)
	router.POST("/task", NewTaskHandler)
	router.GET("/beef/summary", GetBeefSummary)
	err := router.Run()
	if err != nil {
		return
	}
}
