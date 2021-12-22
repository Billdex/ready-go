package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"ready-go/config"
	"time"
)

var DB *gorm.DB

func InitDAO() error {
	dbCfg := config.GetConfig().Mysql
	dao, err := connectDB(dbCfg.DSN, dbCfg.MaxOpenConns, dbCfg.MaxIdleConns, dbCfg.ConnMaxLifetime)
	if err != nil {
		return err
	}
	DB = dao
	return nil
}

func connectDB(dsn string, maxOpenConns int, maxIdleConns int, connMaxLifetime int) (*gorm.DB, error) {
	dao, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
				LogLevel:                  logger.Info,            // Log level:Silent、Error、Warn、Info
				IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,                  // 禁用彩色打印
			},
		),
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
