package guard

import (
	"gin-boilerplate/src/database/dao"
	"gin-boilerplate/src/utils/handler"
	"gin-boilerplate/src/utils/logger"
	"slices"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

func RoleGuard(role ...dao.AuthRole) gin.HandlerFunc {
	return func (c *gin.Context) {
		user, err := handler.GetAuthFromToken(c)

		if err != nil {
			c.AbortWithStatusJSON(handler.Throw500(err.Error()))
			return
		}

		// r, _ := reqRole.(jwt.MapClaims)["role"].(string)
		
		logger.Info("role guard", "role", user.Role)
		if slices.Contains(role, dao.AuthRole(user.Role)) {
			c.Next()
		} else {
			c.AbortWithStatusJSON(handler.Throw500("illegal role request"))
		}
	}
}