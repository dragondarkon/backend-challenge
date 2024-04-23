package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	Assignee string `json:"assignee"`
}

func GetBeefSummary(c *gin.Context) {
	var allBeef = GetAllBeefName()
	var sumBeef = map[string]int{}
	for _, v := range allBeef {
		sumBeef[v] += 1
	}

	c.JSON(http.StatusOK, gin.H{
		"beef": sumBeef,
	})
	c.JSON(http.StatusCreated, sumBeef)
}
