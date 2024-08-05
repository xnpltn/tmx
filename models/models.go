package models

import (
	"gorm.io/gorm"
)

// {"name":"Hello","titles":[{"name":"Tittle 1","dataType":"Text"},{"name":"Title 2","dataType":"Status"}]}
// Sheet represents a spreadsheet with a name, titles, and rows.
type Sheet struct {
	gorm.Model
	Name   string  `json:"name" gorm:"type:varchar(100)"`
	Titles []Title `json:"titles" gorm:"foreignKey:SheetID"`
	Rows   []Row   `json:"rows" gorm:"foreignKey:SheetID"`
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
	DataType       DataType
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
	Name           string   `gorm:"type:text" json:"name"`
	DataType       DataType `gorm:"type:int"`
	DataTypeString string   `json:"dataType" gorm:"type:text"`
	Value          string   `json:"value" gorm:"type:text not null"` // Using text type to store any value.
	RowID          uint     `gorm:"not null"`
}

// testing spreadsheet
var TestSheet = Sheet{
	Name:   "Sheet 2",
	Titles: titles,
	Rows:   []Row{row, row2},
}

var titles = []Title{
	{Name: "title 1"},
	{Name: "title 2"},
	{Name: "title 3"},
	{Name: "title 4"},
	{Name: "title 5"},
	{Name: "title 6"},
}

var row = Row{
	Cells: cells,
}

var row2 = Row{
	Cells: cells2,
}

var cells = []Cell{
	{Value: "1"},
	{Value: "2"},
	{Value: "3"},
	{Value: "4"},
	{Value: "5"},
	{Value: "5"},
}

var cells2 = []Cell{
	{Value: "6"},
	{Value: "7"},
	{Value: "8"},
	{Value: "9"},
	{Value: "10"},
	{Value: "11"},
}
