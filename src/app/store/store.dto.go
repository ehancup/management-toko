package store

import (
	"gin-boilerplate/src/database/dao"
	"gin-boilerplate/src/utils/handler"

	"github.com/gin-gonic/gin"
)

// RESPONSE
type DetailStoreRes struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Avatar      string   `json:"avatar"`
	UserID      uint     `json:"user_id"`
	User        LATRuser `json:"user"`
}
type LATRuser struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func (LATRuser) TableName() string {
	return "auth"
}

type ListAllStoreRes []DetailStoreRes
type ListMyStoreRes []DetailStoreRes


// REQUEST
type CreateStoreReq struct {
	Name        string `json:"name" validate:"required,min=4" example:"store"`
	Description string `json:"description" example:"this is example of desc."`
}

func (v CreateStoreReq) ToEntity(c *gin.Context) dao.StoreEntity {
	user, _ := handler.GetAuthFromToken(c)
	return dao.StoreEntity{
		Name:        v.Name,
		Description: &v.Description,
		UserID:      user.ID,
	}
}

type UpdateStoreReq struct {
	Name        string `json:"name" validate:"required,min=4" example:"store"`
	Description string `json:"description" example:"this is example of desc."`
	Avatar      string `json:"avatar" example:"http://localhost:3000/public/propil.png"`
}

func (v UpdateStoreReq) ToEntity(c *gin.Context) dao.StoreEntity {
	user, _ := handler.GetAuthFromToken(c)
	return dao.StoreEntity{
		Name:        v.Name,
		Description: &v.Description,
		Avatar:      v.Avatar,
		UserID:      user.ID,
	}
}
