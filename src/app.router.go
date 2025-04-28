package src

import (
	"gin-boilerplate/src/app/auth"
	"gin-boilerplate/src/app/product"
	"gin-boilerplate/src/app/store"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	// change if you want to set base path
	route := app.Group("/") 

	auth.InitRoutes(route, GetAuthService())	
	store.InitRoutes(route, GetStoreService())	
	product.InitRoutes(route, GetProductService())
}