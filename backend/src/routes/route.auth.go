package route

import (
	util "mengkodingkan/notes/src/utils"

	"github.com/gin-gonic/gin"
)

func InitAuth(route *gin.Engine) {
	group := route.Group("/api/v1/auth")

	group.POST("/generate", func(c *gin.Context) {
		token, err := util.GenerateToken(map[string]interface{}{"UserBrowser": "chrome"}, "JWT")
		if err != nil {
			c.JSON(400, gin.H{
				"status":  "error",
				"message": "FAILED_TO_GENERATE_TOKEN",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "success",
				"data":   token,
			})
		}
	})
}
