package database

import (
	// "gin-gorm/app/models"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"gin-boilerplate/config"
	"gin-boilerplate/src/database/dao"
	"os"

	// "gin-boilerplate/src/app/users"
	"gin-boilerplate/src/utils/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {

	caCert := os.Getenv("CA_CERT")
	if caCert == "" {
        logger.Fatal("CA_CERT_BASE64 environment variable not set")
    }

    caCertBytes, err := base64.StdEncoding.DecodeString(caCert)
    if err != nil {
        logger.Fatal("Failed to decode CA certificate from base64: %v","data", err)
    }

    rootCertPool := x509.NewCertPool()
    if ok := rootCertPool.AppendCertsFromPEM(caCertBytes); !ok {
        logger.Fatal("Failed to append decoded CA certificate")
    }

    tlsConfig := &tls.Config{
        RootCAs: rootCertPool,
    }

    err = mysqlDriver.RegisterTLSConfig("custom", tlsConfig)
	if err != nil {
        logger.Fatal("Failed to register custom TLS config: %v","data", err)
    }
	var errConnection error
	dbCfg := config.GetConfig().DB
	appCfg := config.GetConfig().App

	DB, errConnection = gorm.Open(mysql.Open(dbCfg.DSN + "&tls=custom"), &gorm.Config{})
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
