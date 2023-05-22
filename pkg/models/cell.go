package models

import "fmt"

type Cell struct {
	ID     string `json:"id"`
	Row    string `json:"row"`
	Column int    `json:"column"`
	Value  string `json:"value"`
}

func NewCell(row string, column int, value string) *Cell {
	return &Cell{
		Row:    row,
		Column: column,
		Value:  value,
	}
}

func (c *Cell) GetPosition() string {
	return c.Row + fmt.Sprint(c.Column)
}
