package presentation

import (
	"log"
	"net/http"

	"github.com/davidzaragoza/tangelo/pkg/domain"
	"github.com/gin-gonic/gin"
)

type Server struct {
	uc *domain.UseCase
}

func NewServer(uc *domain.UseCase) *Server {
	return &Server{uc: uc}
}

func (s *Server) StartServer() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "OK",
			})
		})

		v1.POST("/crop", s.crop)
		v1.GET("/cropped/:name", s.getCroppedImage)

	}

	log.Println("starting server at :8080")
	log.Fatal(router.Run(":8080"))
}

func (s *Server) crop(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("retrieved image %s", header.Filename)
	response, err := s.uc.CropImage(header.Filename, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (s *Server) getCroppedImage(c *gin.Context) {
	name := c.Param("name")
	result, err := s.uc.GetImage(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "image/jpeg", result)
}
