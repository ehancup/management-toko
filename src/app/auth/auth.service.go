package auth

import (
	"errors"
	"gin-boilerplate/src/database"
	"gin-boilerplate/src/database/dao"
	br "gin-boilerplate/src/utils/baseResponse"
	"gin-boilerplate/src/utils/handler"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct{}

// Register godoc
//
//	@Summary		Register
//	@Description	Register with credential
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload		body		RegisterReq	true	"Register Payload"
//	@Success		201			{object}	br.BaseSuccessResponse[any]
//	@Router			/register	[post]
func (Service) Register(c *gin.Context) {
	payload := handler.GetBody[RegisterReq](c)
	if c.IsAborted() {
		return
	}

	entity, errEn := payload.ToEntity()

	if errEn != nil {
		c.JSON(handler.Throw500(errEn.Error()))

		return
	}

	if err := database.DB.Create(&entity).Error; err != nil && strings.HasPrefix(err.Error(), "Error 1062") {
		c.JSON(handler.Throw422("Already account with this email"))
		return
	} else if err != nil {
		c.JSON(handler.Throw500(err.Error()))
		return
	}

	c.JSON(201, br.BaseSuccessResponse[any]{
		Message: "success create account",
	})
}

// Login godoc
// 
//	@Summary		Login
//	@Description	Login 	  with credential
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		LoginReq	true	"login payload"
//	@Success		200		{object}	br.BaseSuccessResponse[any]
//	@Router			/login	[post]
func (Service) Login(c *gin.Context) {
	payload := handler.GetBody[LoginReq](c)
	if c.IsAborted() {
		return
	}
	var user dao.AuthEntity

	if err := database.DB.Where("email = ?", payload.Email).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(handler.Throw404("user not found"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		c.JSON(handler.Throw422("Wrong password!"))
		return
	}

	accessToken, refreshToken, errToken := GenerateToken(user)
	if errToken != nil {
		c.JSON(handler.Throw500(errToken.Error()))
		return
	}

	user.RefreshToken = &refreshToken
	database.DB.Save(&user)

	c.JSON(200, br.BaseSuccessResponse[any]{
		Message: "ok",
		Data: gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}
