package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

// Server struct encapsulates the Gin router and account manager.
type Server struct {
	name   string
	router *gin.Engine
}

// NewServer initializes the server with routes.
func NewServer(name string, logger *zap.Logger) *Server {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		//start := time.Now()
		c.Next()
		//duration := time.Since(start)
		logger.Info("Request",
			//zap.String("method", c.Request.Method),
			//zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			//zap.Duration("duration", duration),
		)
	})

	server := &Server{
		name:   name,
		router: router,
	}
	server.routes()
	return server
}

// routes defines all routes for the server.
func (s *Server) routes() {
	s.router.GET("/name", s.getName)
	s.router.GET("/public-key", s.getPublic)
}

func (s *Server) getPublic(c *gin.Context) {
	publicKey := os.Getenv("PUBLIC_KEY")

	c.JSON(http.StatusOK, gin.H{"publicKey": publicKey})
}

func (s *Server) getName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": s.name})
}

func (s *Server) Run(addr string) {
	fmt.Printf("Server running on http://%s\n", addr)
	go func() {
		if err := s.router.Run(addr); err != nil {
			fmt.Printf("Failed to run server: %v\n", err)
		}
	}()
}
