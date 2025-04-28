package auth

import (
	"fmt"
	"gin-boilerplate/config"
	"gin-boilerplate/src/database/dao"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// REQUEST
type RegisterReq struct {
	Name string `json:"name" validate:"required" example:"tester"`
	Email string `json:"email" validate:"email,required" example:"tester@gmail.com"`
	Password string `json:"password" validate:"required,min=8" example:"12345678"`
}

func (v RegisterReq) ToEntity() (dao.AuthEntity, error) {
	hash, errHash := bcrypt.GenerateFromPassword([]byte(v.Password), 12)

	if errHash != nil {
		
		return dao.AuthEntity{}, errHash
	}

	return dao.AuthEntity{
		Name: v.Name,
		Email: v.Email,
		Password: string(hash),
		Role: dao.User,
	}, nil
}

type LoginReq struct {
	Email string `json:"email" validate:"email,required" example:"tester@gmail.com"`
	Password string `json:"password" validate:"required,min=8" example:"12345678"`
}

// FUNCTION

func GenerateToken(user dao.AuthEntity) ( string, string,  error) {
	
	  
	  at := jwt.NewWithClaims(jwt.SigningMethodHS256, 
	  
	  
		jwt.MapClaims{ 
			"sub": fmt.Sprintf("token%s", strconv.Itoa(int(user.ID))),
			"id" : user.ID,
			"email" : user.Email,
			"role" : user.Role,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
			"iat": time.Now().Unix(),
		})
	  as, err := at.SignedString([]byte(config.GetConfig().App.JwtSecret))
	  if err != nil {
		return "", "", err
	  }
	  rt := jwt.NewWithClaims(jwt.SigningMethodHS256, 
	  
	  
		jwt.MapClaims{ 
			"sub": fmt.Sprintf("token%s", strconv.Itoa(int(user.ID))),
			"id" : user.ID,
			"email" : user.Email,
			"role" : user.Role,
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
			"iat": time.Now().Unix(),
		})
	  rs, err := rt.SignedString([]byte(config.GetConfig().App.JwtSecret))

	  if err != nil {
		return "", "", err
	  }
	  return as,rs, nil
}