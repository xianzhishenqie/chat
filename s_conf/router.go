package s_conf

import "github.com/gin-gonic/gin"

func LoadRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "123",
		})
	})
}
