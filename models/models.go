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

type Cell struct {
	gorm.Model
	SheetID      uint   `gorm:"not null"`
	RowNumber    int    `gorm:"not null"`
	ColumnNumber string `gorm:"size:5;not null"`
	Value        string `gorm:"type:text"`
	DataType     string `gorm:"size:50"`
}
