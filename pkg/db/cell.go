package db

import (
	"fmt"

	"github.com/Deathfireofdoom/excel-client-go/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

func (r *WorkbookRepository) GetCell(id string) (*models.Cell, error) {
	// query the database
	row := r.DB.QueryRow("SELECT ID, row, column, value FROM cells WHERE id = ?", id)

	// parse the result
	var cell models.Cell
	err := row.Scan(&cell.ID, &cell.Row, &cell.Column, &cell.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to get cell: %v", err)
	}

	return &cell, nil
}

func (r *WorkbookRepository) SaveCell(cell *models.Cell, workbookID, sheetID string) error {
	// insert the cell into the database
	_, err := r.DB.Exec(`
	INSERT INTO cells (id, row, column, value, workbook_id, sheet_id)
	VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT (id) DO UPDATE SET
		row = excluded.row,
		column = excluded.column,
		value = excluded.value,
		workbook_id = excluded.workbook_id,
		sheet_id = excluded.sheet_id
	`, cell.ID, cell.Row, cell.Column, cell.Value, workbookID, sheetID)
	if err != nil {
		return err
	}
	return nil
}

func (r *WorkbookRepository) DeleteCell(id string) error {
	// delete the cell from the database
	_, err := r.DB.Exec(`
	DELETE FROM cells
	WHERE id = ?
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete cell: %v", err)
	}
	return nil
}
