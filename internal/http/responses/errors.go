package responses

import "github.com/gin-gonic/gin"

func NotFound(ctx *gin.Context, err error) {
	ctx.JSON(404, gin.H{
		"ok":    false,
		"error": err.Error(),
	})
}

func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(400, gin.H{
		"ok":    false,
		"error": err.Error(),
	})
}

func InternalError(ctx *gin.Context) {
	ctx.JSON(500, gin.H{
		"ok":    false,
		"error": "internal server error",
	})
}
func Unauthorized(ctx *gin.Context) {
	ctx.JSON(401, gin.H{
		"ok":    false,
		"error": "unauthorized",
	})
}
