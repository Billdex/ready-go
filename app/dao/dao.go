package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
	"ready-go/config"
	"ready-go/pkg/logger"
	"time"
)

var DB *gorm.DB

func InitDAO() error {
	dbCfg := config.GetConfig().Mysql
	dao, err := connectDB(dbCfg.DSN, dbCfg.MaxOpenConns, dbCfg.MaxIdleConns, dbCfg.ConnMaxLifetime, zapgorm2.New(logger.L()))
	if err != nil {
		return err
	}
	DB = dao
	return nil
}

func connectDB(dsn string, maxOpenConns int, maxIdleConns int, connMaxLifetime int, log gormlogger.Interface) (*gorm.DB, error) {
	dao, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		Logger:                                   log,
	})
	if err != nil {
		return nil, err
	}
	db, err := dao.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	return dao, nil
}
