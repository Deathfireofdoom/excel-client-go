package db

import (
	"fmt"

	"github.com/Deathfireofdoom/excel-client-go/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

func (r *WorkbookRepository) GetSheet(id string) (*models.Sheet, error) {
	// query the database
	row := r.DB.QueryRow("SELECT ID, pos, name FROM sheets WHERE id = ?", id)

	// parse the result
	var sheet models.Sheet
	err := row.Scan(&sheet.ID, &sheet.Pos, &sheet.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get sheet: %v", err)
	}

	return &sheet, nil
}

func (r *WorkbookRepository) SaveSheet(sheet *models.Sheet) error {
	// insert the sheet into the database
	_, err := r.DB.Exec(`
	INSERT INTO sheets (id, pos, name)
	VALUES (?, ?, ?)
	ON CONFLICT (id) DO UPDATE SET
		pos = excluded.pos,
		name = excluded.name
	`, sheet.ID, sheet.Pos, sheet.Name)
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
