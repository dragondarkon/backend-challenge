package service

import (
	"net/http"
	"piefiredire/external"

	"github.com/gin-gonic/gin"
)

type PieFireDireInterface interface {
	GetBeefSummary(c *gin.Context)
}

type PieFireDireService struct {
	ExternalService external.ExternalInterface
}

func (p *PieFireDireService) GetBeefSummary(c *gin.Context) {
	var allBeef = p.ExternalService.GetAllBeefName()
	var sumBeef = map[string]int{}
	for _, v := range allBeef {
		sumBeef[v] += 1
	}

	c.JSON(http.StatusOK, gin.H{
		"beef": sumBeef,
	})
	c.JSON(http.StatusCreated, sumBeef)
}
