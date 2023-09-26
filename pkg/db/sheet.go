package db

import (
	"database/sql"
	"fmt"

	"github.com/Deathfireofdoom/excel-client-go/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

func (r *WorkbookRepository) GetSheet(id string) (*models.Sheet, error) {
	// query the database for the sheet
	row := r.DB.QueryRow("SELECT ID, pos, name FROM sheets WHERE id = ?", id)

	// parse the sheet result
	var sheet models.Sheet
	err := row.Scan(&sheet.ID, &sheet.Pos, &sheet.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sheet not found for ID: %s", id)
		}
		return nil, fmt.Errorf("failed to get sheet: %v", err)
	}

	// query the database for cells of the sheet
	rows, err := r.DB.Query("SELECT id, row, column, value FROM cells WHERE sheet_id = ?", sheet.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cells for sheet: %v", err)
	}
	defer rows.Close()

	var cells []models.Cell
	for rows.Next() {
		var cell models.Cell
		err := rows.Scan(&cell.ID, &cell.Row, &cell.Column, &cell.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cell: %v", err)
		}
		cells = append(cells, cell)
	}

	// Check for any errors that occurred during the iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during rows iteration: %v", err)
	}

	sheet.Cells = cells

	return &sheet, nil
}

func (r *WorkbookRepository) SaveSheet(sheet *models.Sheet) error {
	// insert the sheet into the database
	_, err := r.DB.Exec(`
	INSERT INTO sheets (workbook_id, id, pos, name)
	VALUES (?, ?, ?, ?)
	ON CONFLICT (id) DO UPDATE SET
		pos = excluded.pos,
		name = excluded.name
	`, sheet.WorkbookID, sheet.ID, sheet.Pos, sheet.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *WorkbookRepository) DeleteSheet(id string) error {
	// delete the sheet from the database
	_, err := r.DB.Exec(`
	DELETE FROM sheets
	WHERE id = ?
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete sheet: %v", err)
	}
	return nil
}
