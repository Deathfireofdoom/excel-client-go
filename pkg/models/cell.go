package models

import (
	"fmt"

	"github.com/Deathfireofdoom/excel-client-go/pkg/utils"
)

type Cell struct {
	ID     string      `json:"id"`
	Row    string      `json:"row"`
	Column int         `json:"column"`
	Value  interface{} `json:"value"`
}

func NewCell(row string, column int, value interface{}, id string) (*Cell, error) {
	// generate uuid
	if id == "" {
		var err error
		id, err = utils.GenerateUUID()
		if err != nil {
			return nil, err
		}
	}

	return &Cell{
		ID:     id,
		Row:    row,
		Column: column,
		Value:  value,
	}, nil
}

func (c *Cell) GetPosition() string {
	return c.Row + fmt.Sprint(c.Column)
}
