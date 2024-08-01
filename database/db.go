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
func ConnectDB(DBURL string) *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open(DBURL), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Println("can't connect to database: ", err.Error())
		os.Exit(-1)
	}
	err = db.AutoMigrate(
		&models.Cell{},
		&models.Sheet{},
		&models.Spreadsheet{},
	)
	if err != nil {
		log.Println("error occured while migrating data: ", err)
		os.Exit(-1)
	}

	log.Println("DB Connected SuccessFully")
	return db
}

// exposes the db
func DB() *gorm.DB {
	// db.Create(&ssheet)
	return db
}

// arbitrary data for testing
var (
	ssheet models.Spreadsheet = models.Spreadsheet{
		Name:   "one",
		Sheets: []models.Sheet{sheet},
	}
	cell models.Cell = models.Cell{
		SheetID:      1,
		RowNumber:    1,
		ColumnNumber: 1,
		DataType:     "int",
		Name:         "name",
	}
	sheet models.Sheet = models.Sheet{
		SpreadsheetID: 1,
		Name:          "one_one",
		Cells:         []models.Cell{cell},
	}
)
