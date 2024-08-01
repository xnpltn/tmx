package models

import (
	"gorm.io/gorm"
)

type Spreadsheet struct {
	gorm.Model
	Name   string  `gorm:"size:255;not null"`
	Sheets []Sheet `gorm:"foreignKey:SpreadsheetID"`
}

type Sheet struct {
	gorm.Model
	SpreadsheetID uint   `gorm:"not null"`
	Name          string `gorm:"size:255;not null"`
	Cells         []Cell `gorm:"foreignKey:SheetID"`
}

/*
example
{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC {0001-01-01 00:00:00 +0000 UTC false}} 2 3 3 hello wolrd  C3}
*/
type Cell struct {
	gorm.Model
	SheetID      uint   `gorm:"not null" json:"sheetID"`
	RowNumber    int    `gorm:"not null" json:"row"`
	ColumnNumber int    `gorm:"size:5;not null" json:"column"`
	Value        string `gorm:"type:text" json:"data"`
	DataType     string `gorm:"size:50"`
	Name         string `gorm:"not null" json:"name"`
}
