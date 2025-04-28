package product

import (
	"gin-boilerplate/src/guard"

	"github.com/gin-gonic/gin"
)

func InitRoutes(c *gin.RouterGroup, s *Service) {
	router := c.Group("/product")
	
	router.POST("/create/:id",guard.AuthGuard(),guard.RoleGuard("user"), s.CreateProduct)
}