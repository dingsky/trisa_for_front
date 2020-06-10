package sqllite

import (
	"os"

	"github.com/trisacrypto/trisa/model/sqlliteModel"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dataBaseSqlLite = "sqlite3"
)

var Database *gorm.DB
var err error

func init() {
	// Connect to database
	Database, err = gorm.Open(dataBaseSqlLite, "./db")
	if err != nil {
		os.Exit(1)
	}

	address := new(sqlliteModel.TblKycList)
	Database.AutoMigrate(address)

	serv := new(sqlliteModel.TblTxnList)
	Database.AutoMigrate(serv)
}

