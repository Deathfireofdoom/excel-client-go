package models

import (
	"fmt"

	"github.com/Deathfireofdoom/excel-client-go/pkg/utils"
)

type Cell struct {
	ID     string      `json:"id"`
	Row    int         `json:"row"`
	Column string      `json:"column"`
	Value  interface{} `json:"value"`
}

func NewCell(row int, column string, value interface{}, id string) (*Cell, error) {
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
	return c.Column + fmt.Sprint(c.Row)
}
