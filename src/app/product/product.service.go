package product

import (
	"errors"
	"gin-boilerplate/src/database"
	"gin-boilerplate/src/database/dao"
	br "gin-boilerplate/src/utils/baseResponse"
	"gin-boilerplate/src/utils/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct{}

// CreateToko	godoc

//	@Summary		Create Product
//	@Description	create toko for 'user' role
//	@Security		BearerAuth
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id						path		uint				true	"id"
//	@Param			payload					body		CreateProductReq	true	"payload"
//	@Success		200						{object}	any
//	@Router			/product/create/{id}	[post]
func (Service) CreateProduct(c *gin.Context) {
	id, errID := handler.CheckID(c.Param("id"))
	if errID != nil {
		c.JSON(handler.Throw500(errID.Error()))
		return
	}

	var store dao.StoreEntity
	if err := database.DB.First(&store, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(handler.Throw(404)("no store found"))
	}

	user, _ := handler.GetAuthFromToken(c)
	if store.UserID != user.ID {
		c.JSON(handler.Throw(500)("youre not the owner"))
		return
	}

	payload := handler.GetBody[CreateProductReq](c)
	if c.IsAborted() {
		return
	}

	en := payload.ToEntity(store.ID)
	if err := database.DB.Create(&en).Error; err != nil {
		c.JSON(handler.Throw(500)(err.Error()))
		return
	}

	c.JSON(201, br.BaseSuccessResponse[any]{
		Message: "success create product",
	})
}
