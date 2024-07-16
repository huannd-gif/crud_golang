package adapters

import (
	"api_crud/core/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type GORMMigrator struct {
	db              *gorm.DB
	autoMigrateList []interface{}
}

func NewGORMMigrator(db *gorm.DB) *GORMMigrator {
	return &GORMMigrator{
		db:              db,
		autoMigrateList: make([]interface{}, 0),
	}
}

func (g *GORMMigrator) addAutoMigrate(model interface{}) {
	g.autoMigrateList = append(g.autoMigrateList, model)
}

func (g *GORMMigrator) MakeMigrations() error {
	err := g.db.AutoMigrate(g.autoMigrateList...)
	return err
}

func NewMysqlConnection(databaseSetting setting.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseSetting.User,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Name,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConnections)
	return db, err
}
