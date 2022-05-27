package response

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, map[string]interface{}{
		"data":       data,
		"error_no":   code,
		"error_msg":  msg,
	})
}

func AbortResponse(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.AbortWithStatusJSON(httpCode, map[string]interface{}{
		"data":       data,
		"error_no":   httpCode,
		"error_msg":  msg,
	})
}

func ClientErr(c *gin.Context, msg string) {
	Response(c, 400, msg, nil)
}

func ServerErr(c *gin.Context, msg string) {
	Response(c, 500, msg, nil)
}

func ServerSuccess(c *gin.Context, data interface{}) {
	Response(c, 0, "success", data)
}

func Success(c *gin.Context) {
	Response(c, 0, "success", nil)
}

