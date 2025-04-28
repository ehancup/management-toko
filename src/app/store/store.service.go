package store

import (
	"errors"
	"gin-boilerplate/src/database"
	"gin-boilerplate/src/database/dao"
	br "gin-boilerplate/src/utils/baseResponse"
	"gin-boilerplate/src/utils/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type Service struct{}

// GetAllToko	godoc

//	@Summary		get all toko (admin)
//	@Description	get all toko with pagination for admin
//	@Security		BearerAuth
//	@Tags			store
//	@Accept			json
//	@Produce		json
//	@Param			page			query		int													false	"page"
//	@Param			page_size		query		int													false	"page size"
//	@Param			q				query		string												false	"search"
//	@Success		200				{object}	br.BaseSuccessResponsePagination[ListAllStoreRes]	
//	@Router			/store/list-all	[get]
func (Service) GetAllStore(c *gin.Context) {

	pg, errPg := handler.GetPagination(c)
	if errPg != nil {
		c.JSON(handler.Throw500(errPg.Error()))
		return
	}

	var tokos ListAllStoreRes
	var totalTokos int64

	tx := database.DB.Model(&dao.StoreEntity{}).Preload("User")
	if q := c.Query("q"); q != "" {
		tx.Where("name LIKE ?", "%"+q+"%")
	}
	tx.Count(&totalTokos)
	tx.Offset((pg.Page - 1) * pg.PageSize).Limit(pg.PageSize)
	if err := tx.Find(&tokos).Error; err!= nil {
		c.JSON(handler.Throw500(err.Error()))
		return
	}
	c.JSON(200, br.BaseSuccessResponsePagination[ListAllStoreRes]{
		Message: "success",
		Data: tokos,
		Pagination: pg.GetResponse(int(totalTokos)),
	})
}


// CreateToko	godoc

//	@Summary		Create Toko (user)
//	@Description	create toko for 'user' role
//	@Security		BearerAuth
//	@Tags			store
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		CreateStoreReq				true	"payload"
//	@Success		201				{object}	br.BaseSuccessResponse[any]	
//	@Router			/store/create	[post]
func (Service) CreateStore(c *gin.Context) {
	payload := handler.GetBody[CreateStoreReq](c)
	if c.IsAborted() {
		return
	}
	en := payload.ToEntity(c)
	if err := database.DB.Create(&en).Error; err != nil {
		c.JSON(handler.Throw500(err.Error()))
		return
	}
	c.JSON(201, br.BaseSuccessResponse[any]{
		Message: "success",
	})
} 

// GetMyToko	godoc

//	@Summary		get my toko (user)
//	@Description	get my toko with pagination for user
//	@Security		BearerAuth
//	@Tags			store
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int													false	"page"
//	@Param			page_size	query		int													false	"page size"
//	@Success		200			{object}	br.BaseSuccessResponsePagination[ListMyStoreRes]	
//	@Router			/store/my	[get]
func (Service) GetMyStore(c *gin.Context) {
	pg, errPg := handler.GetPagination(c)
	if errPg != nil {
		c.JSON(handler.Throw500(errPg.Error()))
		return
	}

	var tokos ListMyStoreRes
	var totalTokos int64

	user, _ := handler.GetAuthFromToken(c)

	tx := database.DB.Model(&dao.StoreEntity{}).Preload("User")
	tx.Where("user_id = ?", user.ID)
	tx.Count(&totalTokos)
	tx.Offset((pg.Page - 1) * pg.PageSize).Limit(pg.PageSize)
	if err := tx.Find(&tokos).Error; err!= nil {
		c.JSON(handler.Throw500(err.Error()))
		return
	}
	c.JSON(200, br.BaseSuccessResponsePagination[ListMyStoreRes]{
		Message: "success",
		Data: tokos,
		Pagination: pg.GetResponse(int(totalTokos)),
	})
}

// UpdateToko godoc

//	@Summary		update (owner)
//	@Description	update store 
//	@Tags			store
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id					path		int				true	"id"
//	@Param			payload				body		UpdateStoreReq	true	"payload"
//	@Success		200					{object}	br.BaseSuccessResponse[any]
//	@Router			/store/update/{id}	[put]
func (Service) UpdateStore(c *gin.Context) {
	id, errId := handler.CheckID(c.Param("id"))

	if errId != nil {
		c.JSON(handler.Throw500(errId.Error()))
		return
	}

	var toko dao.StoreEntity

	if err := database.DB.Where("id = ?", id).First(&toko).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(handler.Throw404("toko not found"))
		return
	}

	user, _ := handler.GetAuthFromToken(c)
	if toko.UserID != user.ID {
		c.JSON(handler.Throw500("youre not the owner"))
		return
	}

	payload := handler.GetBody[UpdateStoreReq](c)
	if c.IsAborted() {
		return
	}
	en := payload.ToEntity(c)
	en.ID = id

	if err := database.DB.Save(&en).Error; err != nil {
		c.JSON(handler.Throw500(err.Error()))
		return
	}

	c.JSON(200, br.BaseSuccessResponse[any]{
		Message: "success update data",
	})
}

// DetailToko godoc

//	@Summary		Detail
//	@Description	Get detail store
//	@Tags			store
//	@Accept			json
//	@Produce		json
//	@Param			id					path		int	true	"id"
//	@Success		200					{object}	br.BaseSuccessResponse[DetailStoreRes]
//	@Router			/store/detail/{id}	[get]
func (Service) GetDetailStore(c *gin.Context) {
	id, errID := handler.CheckID(c.Param("id"))
	if errID != nil {
		c.JSON(handler.Throw500(errID.Error()))
		return
	}

	var store DetailStoreRes

	if err := database.DB.Model(&dao.StoreEntity{}).Preload("User").First(&store, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(handler.Throw404("no store found"))
		return
	}

	c.JSON(200, br.BaseSuccessResponse[DetailStoreRes]{
		Message: "success",
		Data: store,
	})
}