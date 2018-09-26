package persist

import (
	"github.com/jinzhu/gorm"

	"github.com/Eric-GreenComb/one-pushinfo/bean"
	"github.com/Eric-GreenComb/one-pushinfo/config"
)

// ConnectDb connect Db
func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open(config.MariaDB.Dialect, config.MariaDB.URL)

	if config.Server.GormLogMode == "false" {
		db.LogMode(false)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

// InitDatabase Init Db
func InitDatabase() {
	db, err := gorm.Open(config.MariaDB.Dialect, config.MariaDB.URL)
	defer db.Close()

	if config.Server.GormLogMode == "false" {
		db.LogMode(false)
	}

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&bean.Order{}) {
		db.CreateTable(&bean.Order{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.Order{})
	}

	return
}
