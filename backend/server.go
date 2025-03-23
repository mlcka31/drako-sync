package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// Server struct encapsulates the Gin router and account manager.
type Server struct {
	name      string
	router    *gin.Engine
	publicKey string
}

// NewServer initializes the server with routes.
func NewServer(name string, logger *zap.Logger) *Server {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Next()
		logger.Info("Request",
			zap.Int("status", c.Writer.Status()),
		)
	})

	server := &Server{
		name:   name,
		router: router,
	}
	server.routes()
	return server
}

// SetPublicKey
func (s *Server) SetPublicKey(publicKey string) {
	s.publicKey = publicKey
}

// routes defines all routes for the server.
func (s *Server) routes() {
	s.router.GET("/name", s.getName)
	s.router.GET("/public-key", s.getPublic)
}

func (s *Server) getPublic(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"publicKey": s.publicKey})
}

func (s *Server) getName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": s.name})
}

func (s *Server) Run(addr string) {
	fmt.Printf("To get the private key go to http://%s/private-key\n", addr)
	go func() {
		if err := s.router.Run(addr); err != nil {
			fmt.Printf("Failed to run server: %v\n", err)
		}
	}()
}
