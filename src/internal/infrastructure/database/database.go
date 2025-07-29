package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"notion/src/config"
)

// Database struct chứa database connection
type Database struct {
	DB *gorm.DB
}

// NewDatabase tạo connection mới đến database
func NewDatabase(cfg config.IConfig) *Database {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.GetDBUser(),
		cfg.GetDBPassword(),
		cfg.GetDBHost(),
		cfg.GetDBPort(),
		cfg.GetDBName(),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate các bảng
	//err = db.AutoMigrate(&domain.User{})
	//if err != nil {
	//	log.Fatal("Failed to migrate database:", err)
	//}

	return &Database{DB: db}
}

// GetDB trả về database instance
func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
