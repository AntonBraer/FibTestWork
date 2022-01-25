package api

import (
	"fbsTest/config"
	"fbsTest/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config  config.Config
	router  *gin.Engine
	service *service.FibService
}

// NewServer создает сервер
func NewServer(service *service.FibService, config config.Config) *Server {
	return &Server{
		service: service,
		config:  config,
	}
}

// Run запускает сервер
func (s *Server) Run(url string) error {
	return s.router.Run(url)
}
