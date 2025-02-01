package routes

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"image/color"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
	"sdcraft.fun/oauth2/database"
	"sdcraft.fun/oauth2/globals"
	"sdcraft.fun/oauth2/models"
	"sdcraft.fun/oauth2/utils"
)

type captchaPayload struct {
	ChallengeID string `json:"challenge_id"`
	Answer      string `json:"answer"`
}

type registerRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

var (
	emailRegex = regexp.MustCompile(`^\w+([-+.]?\w+)*@\w+([-.]?\w+)*\.\w+([-.]?\w+)*$`)
	//passwordRegex = regexp.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,18}$`)
	nameRegex      = regexp.MustCompile(`^[a-zA-Z0-9_]{5,20}$`)
	captchaStroage *utils.CaptchaCache
	BgColor        = color.RGBA{R: 255, G: 255, B: 255, A: 255}
)

func init() {
	var err error
	captchaStroage, err = utils.NewCaptchaCache()
	if err != nil {
		logrus.Fatalf("Failed to create captcha stroage: %v", err)
	}
}

func Register_v1_routes(g *gin.RouterGroup) {
	g.GET("/publicKey", func(ctx *gin.Context) {
		ctx.JSON(globals.Success, gin.H{
			"code": globals.Success,
			"key":  globals.RSAPublicKey,
		})
	})
	g.GET("/captcha", func(ctx *gin.Context) {
		driver := base64Captcha.NewDriverMath(
			40,
			120,
			10,
			5,
			&BgColor,
			utils.DefaultEmbeddedFonts,
			[]string{"JetBrainsMono-Bold.ttf"},
		)
		c := base64Captcha.NewCaptcha(driver, captchaStroage)
		id, b64s, _, err := c.Generate()
		if err != nil {
			ctx.JSON(500, gin.H{
				"code":    globals.GenerateCaptchaFailed,
				"message": "Failed to generate captcha.",
			})
			return
		}
		ctx.JSON(globals.Success, gin.H{
			"code":   globals.Success,
			"id":     id,
			"base64": b64s,
		})
	})
	g.POST("/register", CaptchaVerifyMiddleware, registerEndpoint)
	g.POST("/login", CaptchaVerifyMiddleware)
}

func CaptchaVerifyMiddleware(c *gin.Context) {
	var req struct {
		Verify captchaPayload `json:"verify"`
	}
	if err := c.MustBindWith(&req, binding.JSON); err != nil {
		c.JSON(400, gin.H{
			"code":    globals.InvalidRequest,
			"message": "Invalid request.",
		})
		c.Abort()
		return
	}
	if !captchaStroage.Verify(req.Verify.ChallengeID, req.Verify.Answer, true) {
		c.JSON(400, gin.H{
			"code":    globals.InvalidCaptcha,
			"message": "Invalid captcha.",
		})
		c.Abort()
		return
	}
	c.Next()
}

func registerEndpoint(ctx *gin.Context) {
	var req registerRequest
	//TODO Captcha Verify
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":    globals.InvalidRequest,
			"message": "Invalid request.",
		})
		return
	}

	if !emailRegex.MatchString(req.Email) {
		ctx.JSON(400, gin.H{
			"code":    globals.InvalidEmail,
			"message": "Invalid email.",
		})
		return
	}

	if !nameRegex.MatchString(req.Username) {
		ctx.JSON(400, gin.H{
			"code":    globals.InvalidUsername,
			"message": "Invalid username.",
		})
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(globals.RSAPrivateKey)
	if err != nil {
		logrus.Errorf("Unexpected Exception occurred when parse RSA prikey, cause: %v", err)
		ctx.JSON(500, gin.H{
			"code":    globals.ParseRSAPrivateKeyError,
			"message": "Unexpected Exception occurred when parse RSA prikey.",
		})
		return
	}

	encryptedData, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    globals.DecodePasswordError,
			"message": "Unexpected Exception occurred when decode password.",
		})
		return
	}

	binaryPassword, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    globals.DecodePasswordError,
			"message": "Unexpected Exception occurred when decode password.",
		})
		return
	}

	var pw = string(binaryPassword)
	if !utils.ValidatePassword(pw) {
		ctx.JSON(400, gin.H{
			"code":    globals.InvalidPassword,
			"message": "Invaild password.",
		})
		return
	}

	u := models.NewUser(req.Username, req.Email, utils.HashPassword(pw, []byte(globals.Generate.SALT)))
	tx := database.DB.Create(u)
	if tx.Error != nil {
		ctx.JSON(400, gin.H{
			"code":    globals.CreateUserFailed,
			"message": "Failed to create user. Is the email or username used?",
		})
		return
	}

	token, err := utils.Sign(u.ID, privateKey)
	if err != nil {
		logrus.Errorf("Unexpected Exception occurred when sign token, cause: %v", err)
		ctx.JSON(500, gin.H{
			"code":    globals.SignTokenError,
			"message": "Unexpected Exception occurred when sign token.",
		})
		return
	}

	ctx.JSON(globals.Success, gin.H{
		"code":    globals.Success,
		"message": "Success!",
		"token":   token,
	})
}
