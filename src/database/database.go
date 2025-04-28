package database

import (
	// "gin-gorm/app/models"
	"gin-boilerplate/config"
	"gin-boilerplate/src/database/dao"

	// "gin-boilerplate/src/app/users"
	"gin-boilerplate/src/utils/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var errConnection error
	dbCfg := config.GetConfig().DB
	appCfg := config.GetConfig().App

	DB, errConnection = gorm.Open(mysql.Open(dbCfg.DSN), &gorm.Config{})
	if (appCfg.Mode == "debug") {DB.Debug()}

	if errConnection != nil {
		logger.Fatal("Cant connect to database")
	}

	logger.Info("Success connected to database")

}

func tableList() []any {
	return []any{
		&dao.AuthEntity{},
		&dao.StoreEntity{},
		&dao.ProductEntity{},
	}
}

func Migrate(reset bool) {
	if reset {
		DB.Migrator().DropTable(tableList()...)
	}
	err := DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(tableList()...)

	if err != nil {
		logger.Fatal("Migration Failed", "err", err.Error())
	}
	logger.Info("Database migrated successfully.")
}
