package store

import (
	g "gin-boilerplate/src/guard"

	"github.com/gin-gonic/gin"
)

func InitRoutes(c *gin.RouterGroup, s *Service) {
	router := c.Group("/store")

	router.GET("/list-all", g.AuthGuard(), g.RoleGuard("admin"), s.GetAllStore)
	router.GET("/my", g.AuthGuard(), g.RoleGuard("user"), s.GetMyStore)
	router.POST("/create", g.AuthGuard(), g.RoleGuard("user"), s.CreateStore)
	router.PUT("/update/:id", g.AuthGuard(), g.RoleGuard("user"), s.UpdateStore)
	router.GET("/detail/:id", s.GetDetailStore)
}