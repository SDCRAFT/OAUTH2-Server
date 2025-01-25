package routes

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sdcraft.fun/oauth2/globals"
)

type registerRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

var emailRegex = regexp.MustCompile(`^\w+([-+.]?\w+)*@\w+([-.]?\w+)*\.\w+([-.]?\w+)*$`)

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
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		logrus.Print(err.Error())
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

	privateKey, err := x509.ParsePKCS1PrivateKey(globals.RSAPrivateKey)

	if err != nil {
		logrus.Errorf("Unexpected Exception occurred when parse RSA prikey, cause: %v", err)
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "Unexpected Exception occurred when parse RSA prikey",
		})
		return
	}

	encryptedData, err := base64.StdEncoding.DecodeString(req.Password)

	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Unexpected Exception occurred when decode password",
		})
		return
	}

	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		logrus.Errorf("%v", err)
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Unexpected Exception occurred when decode password",
		})
		return
	}
	if string(decryptedData) == "1" {

	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "Success",
	})
}
