// package database contains all database functionality
package database

import (
	"log"
	"os"

	"spreadsheet/spdb/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// creates a connection to sqlite database
func ConnectDB(DBURL string) {
	var err error
	db, err = gorm.Open(sqlite.Open(DBURL), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Println("can't connect to database: ", err.Error())
		os.Exit(-1)
	}
	err = db.AutoMigrate(
		&models.Sheet{},
		&models.Cell{},
		&models.Row{},
		&models.Title{},
	)
	if err != nil {
		log.Println("error occured while migrating data: ", err)
		os.Exit(-1)
	}
	log.Println("DB Connected SuccessFully")
}

// exposes the db
func DB() *gorm.DB {
	return db
}
