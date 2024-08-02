package models

import (
	"gorm.io/gorm"
)

// {"name":"Hello","titles":[{"name":"Tittle 1","dataType":"Text"},{"name":"Title 2","dataType":"Status"}]}
// Sheet represents a spreadsheet with a name, titles, and rows.
type Sheet struct {
	gorm.Model
	Name    string  `json:"name" gorm:"type:varchar(100)"`
	Tittles []Title `json:"titles" gorm:"foreignKey:SheetID"`
	Rows    []Row   `json:"rows" gorm:"foreignKey:SheetID"`
}

// Row represents a row in a spreadsheet, containing multiple cells.
type Row struct {
	gorm.Model
	Cells   []Cell `json:"cells" gorm:"foreignKey:RowID"`
	SheetID uint   `json:"sheet_id" gorm:"not null"`
}

type Title struct {
	gorm.Model
	Name           string `json:"name"`
	DataType       string
	DataTypeString string `json:"dataType"`
	SheetID        uint   `gorm:"not null"`
}

// DataType represents the type of data a cell holds.
type DataType uint8

const (
	Text DataType = iota
	CheckBox
	Status
	Tag
	Date
	Number
	Label
)

// Cell represents a single cell in a row, holding a specific type of data.
type Cell struct {
	gorm.Model
	DataType       DataType    `gorm:"type:int"`
	DataTypeString string      `json:"dataType" gorm:"type:text"`
	Value          interface{} `json:"value" gorm:"type:text"` // Using text type to store any value.
	RowID          uint        `json:"row_id" gorm:"not null"`
}
