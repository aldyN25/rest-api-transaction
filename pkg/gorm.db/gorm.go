package gormdb

import (
	"fmt"
	"sync"

	"github.com/aldyN25/myproject/app/config"
	"github.com/aldyN25/myproject/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}
var db *gorm.DB

func GetInstance() (*gorm.DB, error) {
	if db == nil {
		config := config.GetInstance()

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			config.Dbconfig.Host,
			config.Dbconfig.Username,
			config.Dbconfig.Password,
			config.Dbconfig.Dbname,
			config.Dbconfig.Port,
		)
		lock.Lock()
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		lock.Unlock()
		if err != nil {
			return nil, err
		}

		if config.Dbconfig.DbIsMigrate {
			// Migrate Here
			db.AutoMigrate(&models.Users{}, &models.Transactions{})
		}
		return db, nil
	}

	return db, nil
}
