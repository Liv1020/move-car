package components

import "github.com/gin-gonic/gin"

// ResponseSuccess ResponseSuccess
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "ok",
		"data":    data,
	})
}

// ResponseError ResponseError
func ResponseError(c *gin.Context, code int, err error) {
	c.JSON(200, gin.H{
		"status":  code,
		"message": err,
	})
}
