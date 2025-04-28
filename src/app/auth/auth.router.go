package auth

import "github.com/gin-gonic/gin"

func InitRoutes(c *gin.RouterGroup, s *Service) {
	c.POST("/register", s.Register)
	c.POST("/login", s.Login)
}