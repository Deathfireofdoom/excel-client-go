package models

import "github.com/Deathfireofdoom/excel-client-go/pkg/utils"

type Sheet struct {
	ID string `json:"id"`

	Pos   int    `json:"pos"`
	Name  string `json:"name"`
	Cells []Cell `json:"cells"`
}

func NewSheet(pos int, name, id string) (*Sheet, error) {
	// generate uuid
	if id == "" {
		var err error
		id, err = utils.GenerateUUID()
		if err != nil {
			return nil, err
		}
	}

	return &Sheet{
		ID:    id,
		Pos:   pos,
		Name:  name,
		Cells: []Cell{},
	}, nil

}
