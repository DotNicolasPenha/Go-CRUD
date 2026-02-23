package responses

import "github.com/gin-gonic/gin"

func NotFound(ctx *gin.Context, msg string) {
	ctx.JSON(404, gin.H{
		"ok":    false,
		"error": msg,
	})
}

func BadRequest(ctx *gin.Context, msg string) {
	ctx.JSON(400, gin.H{
		"ok":    false,
		"error": msg,
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
