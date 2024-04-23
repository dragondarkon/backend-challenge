package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"piefiredire/service"
)

// Router to apis lisening
type PieFireDireRouter struct {
	PieFireDireHandler service.PieFireDireService
}

// CORSMiddleware is Cross-Origin Resource Sharing Middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func health(c *gin.Context) {

	type Resp struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	c.JSON(http.StatusOK, Resp{
		Code:    "200",
		Message: "Version 0.0.1 OK",
	})

}

// Route is setup router
func (router PieFireDireRouter) Route() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/health", health)

	r.GET("/beef/summary", router.PieFireDireHandler.GetBeefSummary)

	return r
}
