package dbase

import (
	"time"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/go-rest-template/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
}

/*
Create connection pooling using gorm postgres driver.
*/
func PgConnect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(app.Config.Database.Postgres.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.New().Error("Error connecting Postgres database: %v", err)
		panic(err)
	}

	// connection pooling configuration
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(app.Config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(app.Config.Database.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(app.Config.Database.MaxIdleTime) * time.Second)

	return db
}
