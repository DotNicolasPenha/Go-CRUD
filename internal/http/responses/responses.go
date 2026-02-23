package responses

import "github.com/gin-gonic/gin"

func OK_DATA(ctx *gin.Context, data any) {
	ctx.JSON(200, gin.H{
		"ok":   true,
		"data": data,
	})
}
func OK_CREATED(ctx *gin.Context, msg string) {
	ctx.JSON(201, gin.H{
		"ok":  true,
		"msg": msg,
	})
}
func OK_DELETE(ctx *gin.Context) {
	ctx.Status(204)
}
