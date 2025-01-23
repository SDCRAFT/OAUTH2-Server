package routes

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"sdcraft.fun/oauth2/globals"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var emailRegex = regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w{2,}$`)

func Register_v1_routes(g *gin.RouterGroup) {
	g.GET("/publicKey", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 200,
			"key":  globals.RSAPublicKey,
		})
	})
	g.POST("/register", registerEndpoint)
}

func registerEndpoint(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":  400,
			"error": "Invalid request",
		})
		return
	}
	if !emailRegex.MatchString(req.Email) {
		ctx.JSON(400, gin.H{
			"code":  400,
			"error": "Invalid email",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "Succeeded",
	})
}
