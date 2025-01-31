package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	database "sdcraft.fun/oauth2/database"
	"sdcraft.fun/oauth2/globals"
	Routes "sdcraft.fun/oauth2/routes"
	"sdcraft.fun/oauth2/utils"
)

var v = viper.New()

func readConfig() {
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("json")
	v.SetDefault("Config", &globals.Config)
	v.SetDefault("Generate", &globals.Generate)
	err := v.ReadInConfig()
	if err != nil {
		e := v.SafeWriteConfig()
		if e != nil {
			logrus.Fatalf("Failed to read config file: %v", err)
			logrus.Fatalf("Failed to save config file: %v", e)
		}

	}
	err = v.UnmarshalKey("Config", &globals.Config)
	if err != nil {
		logrus.Fatalf("Failed to unmarshal config file: %v", err)
	}
	err = v.UnmarshalKey("Generate", &globals.Generate)
	if err != nil {
		logrus.Fatalf("Failed to unmarshal config file: %v", err)
	}
}

type PrettyFormatter struct{}

func (f *PrettyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.UTC().Format("2006-01-02 15:04:05")
	var msg string
	if entry.Data["name"] != nil {
		msg = fmt.Sprintf("[%s][%s][%s] %s\n", timestamp, entry.Data["name"], entry.Level, entry.Message)
	} else {
		msg = fmt.Sprintf("[%s][%s] %s\n", timestamp, entry.Level, entry.Message)
	}
	return []byte(msg), nil
}

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&PrettyFormatter{})
	logrus.Info("Starting...")
	gin.SetMode(gin.ReleaseMode)
	readConfig()
	defer viper.SafeWriteConfig()
	database.Init()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(LoggerMiddleware())
	Routes.Register_v1_routes(router.Group("/api/v1"))
	logrus.Infof("Server will listen on: %s:%d", globals.Config.Listen.Host, globals.Config.Listen.Port)
	err := router.Run(fmt.Sprintf("%s:%d", globals.Config.Listen.Host, globals.Config.Listen.Port))
	if err != nil {
		logrus.Fatalf("Failed listen the address: %v", err)
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logrus.WithFields(logrus.Fields{"name": "GIN"}).
			Info(utils.Map2String(map[string]interface{}{
				"code":     statusCode,
				"duration": latencyTime,
				"ip":       clientIP,
				"method":   reqMethod,
				"uri":      reqUri,
			}, " "))
	}
}
