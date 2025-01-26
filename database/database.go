package database

import (
	"context"
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
			}))
	} else {
		logrus.WithFields(logrus.Fields{"name": "GORM"}).Trace(
			utils.Map2String(map[string]interface{}{
				"message":  "SQL query executed",
				"sql":      sql,
				"rows":     rows,
				"duration": elapsed,
			}))
	}
}

func Init() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		Logger: &LogrusLogger{},
	})

	if err != nil {
		logrus.Fatalf("Failed to connect database: %v", err)
	}
	DB = db
	DB.AutoMigrate(&Models.User{})
}
