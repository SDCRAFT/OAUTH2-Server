package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sdcraft.fun/oauth2/globals"
	"sdcraft.fun/oauth2/models"
	Models "sdcraft.fun/oauth2/models"
	"sdcraft.fun/oauth2/utils"
)

var DB *gorm.DB

type LogrusLogger struct {
}

func (l *LogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *LogrusLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithFields(logrus.Fields{"name": "GORM"}).WithContext(ctx).Info("message=" + fmt.Sprintf(msg, data...))
}

func (l *LogrusLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithFields(logrus.Fields{"name": "GORM"}).WithContext(ctx).Warn("message=" + fmt.Sprintf(msg, data...))
}

func (l *LogrusLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithFields(logrus.Fields{"name": "GORM"}).WithContext(ctx).Error("message=" + fmt.Sprintf(msg, data...))
}

func (l *LogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)
	if err != nil {
		logrus.WithFields(logrus.Fields{"name": "GORM"}).Trace(
			utils.Map2String(map[string]interface{}{
				"message":  "SQL query failed",
				"sql":      sql,
				"rows":     rows,
				"duration": elapsed,
			}, " "))
	} else {
		logrus.WithFields(logrus.Fields{"name": "GORM"}).Trace(
			utils.Map2String(map[string]interface{}{
				"message":  "SQL query executed",
				"sql":      sql,
				"rows":     rows,
				"duration": elapsed,
			}, " "))
	}
}

func dsnBuilder(db models.Database) string {
	switch strings.ToLower(db.Type) {
	case "sqlite":
		{
			return db.Database
		}
	case "mysql":
		{
			convertedMap := make(map[string]interface{})
			for key, value := range db.Paramters {
				convertedMap[key] = value
			}
			return fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?%s",
				db.Account.Username,
				db.Account.Password,
				db.Host,
				db.Port,
				db.Database,
				utils.Map2String(convertedMap, ";"),
			)
		}
	default:
		{
			logrus.Fatalf("Invaild database type: %s", db.Type)
			return ""
		}
	}
}

func Init() {
	db, err := gorm.Open(sqlite.Open(dsnBuilder(globals.Config.Database)), &gorm.Config{
		Logger: &LogrusLogger{},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: globals.Config.Database.TablePrefix,
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to connect database: %v", err)
	}
	DB = db
	DB.AutoMigrate(&Models.User{})
}
