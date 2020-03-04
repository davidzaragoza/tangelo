package presentation

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (p *Server) StartServer() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "OK",
			})
		})
	}

	log.Println("starting server at :8080")
	log.Fatal(router.Run(":8080"))
}
