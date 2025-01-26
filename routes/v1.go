package routes

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sdcraft.fun/oauth2/database"
	"sdcraft.fun/oauth2/globals"
	"sdcraft.fun/oauth2/models"
	"sdcraft.fun/oauth2/utils"
)

type registerRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

var (
	emailRegex = regexp.MustCompile(`^\w+([-+.]?\w+)*@\w+([-.]?\w+)*\.\w+([-.]?\w+)*$`)
	//passwordRegex = regexp.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,18}$`)
	nameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{5,20}$`)
)

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
		ctx.JSON(400, gin.H{
			"code":    -401,
			"message": "Invalid request.",
		})
		return
	}
	if !emailRegex.MatchString(req.Email) {
		ctx.JSON(400, gin.H{
			"code":    -402,
			"message": "Invalid email.",
		})
		return
	}

	if !nameRegex.MatchString(req.Username) {
		ctx.JSON(400, gin.H{
			"code":    -403,
			"message": "Invalid username.",
		})
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(globals.RSAPrivateKey)

	if err != nil {
		logrus.Errorf("Unexpected Exception occurred when parse RSA prikey, cause: %v", err)
		ctx.JSON(500, gin.H{
			"code":    -500,
			"message": "Unexpected Exception occurred when parse RSA prikey.",
		})
		return
	}

	encryptedData, err := base64.StdEncoding.DecodeString(req.Password)

	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    -404,
			"message": "Unexpected Exception occurred when decode password.",
		})
		return
	}

	binaryPassword, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    -404,
			"message": "Unexpected Exception occurred when decode password.",
		})
		return
	}
	var pw = string(binaryPassword)
	if !utils.ValidatePassword(pw) {
		ctx.JSON(400, gin.H{
			"code":    -405,
			"message": "Invaild password.",
		})
		return
	}
	tx := database.DB.Create(models.NewUser(req.Username, req.Email, utils.HashPassword(pw, []byte(globals.Generate.SALT))))
	if tx.Error != nil {
		ctx.JSON(400, gin.H{
			"code":    -406,
			"message": "Failed to create user. Is the email or username used?",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "Success!",
	})
}
