package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "mengkodingkan/notes/src/utils"
)

type UnathorizedError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
}

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var errResponse UnathorizedError

		errResponse.Status = "Forbidden"
		errResponse.Message = "Unauthorized, please put your token in your header"
		errResponse.Code = http.StatusForbidden
		errResponse.Method = c.Request.Method

		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusForbidden, errResponse)
			defer c.AbortWithStatus(http.StatusForbidden)
		} else {
			token, err := util.VerifyTokenHeader(c, "JWT")

			errResponse.Status = "Unathorizated"
			errResponse.Code = http.StatusUnauthorized
			errResponse.Method = c.Request.Method
			errResponse.Message = "accessToken invalid or expired"

			if err != nil {
				c.JSON(http.StatusUnauthorized, errResponse)
				defer c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				c.Set("user", token.Claims)
				c.Next()
			}

		}
	})
}
