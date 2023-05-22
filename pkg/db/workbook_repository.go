package db

import (
	"database/sql"
	"fmt"

	"github.com/Deathfireofdoom/excel-client-go/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

type WorkbookRepository struct {
	DB *SQLiteDB
}

func NewWorkbookRepository() (*WorkbookRepository, error) {
	// initialize sql-lite database
	db, err := NewSQLiteDB("excel.db")
	if err != nil {
		return nil, err
	}

	// initialize repository
	repository := &WorkbookRepository{DB: db}
	err = repository.Initialize()
	if err != nil {
		return nil, err
	}

	return &WorkbookRepository{DB: db}, nil
}

func (r *WorkbookRepository) Initialize() error {
	// create table for workbooks
	_, err := r.DB.Exec(`
		CREATE TABLE IF NOT EXISTS workbooks (
			id              TEXT PRIMARY KEY UNIQUE,
			file_name       TEXT,
			extension       TEXT,
			folder_path     TEXT,
			last_modified   DATETIME DEFAULT CURRENT_TIMESTAMP
			)`)
	if err != nil {
		return err
	}

	// create table for sheets
	_, err = r.DB.Exec(`
		CREATE TABLE IF NOT EXISTS sheets (
			id              TEXT PRIMARY KEY,
			workbook_id     TEXT,
			pos             INTEGER,
			name            TEXT,
			last_updated_at TIMESTAMP DEAFULT CURRENT_TIMESTAMP,
			FOREIGN KEY (workbook_id) REFERENCES workbooks(id)
			)`)
	if err != nil {
		return err
	}

	// create table for cells
	_, err = r.DB.Exec(`
		CREATE TABLE IF NOT EXISTS cells (
			id              TEXT PRIMARY KEY,
			sheet_id        TEXT,
			workbook_id     TEXT,
			row             TEXT,
			column          INTEGER,
			value           TEXT,
			last_updated_at TIMESTAMP DEAFULT CURRENT_TIMESTAMP,
			FOREIGN KEY (sheet_id) REFERENCES sheets(id),
			FOREIGN KEY (workbook_id) REFERENCES workbooks(id)
		)`)
	if err != nil {
		return err
	}

	return nil
}

func (r *WorkbookRepository) SaveMetadata(workbook *models.Workbook) error {
	_, err := r.DB.Exec(`
	INSERT INTO workbooks (id, file_name, extension, folder_path)
	VALUES (?, ?, ?, ?)
	ON CONFLICT (id) DO UPDATE SET
		file_name = excluded.file_name,
		extension = excluded.extension,
		folder_path = excluded.folder_path
	`, workbook.ID, workbook.FileName, workbook.Extension, workbook.FolderPath)
	if err != nil {
		return err
	}
	return nil
}

func (r *WorkbookRepository) GetMetadata(id string) (*models.Workbook, error) {
	var workbook models.Workbook
	row := r.DB.QueryRow(`
	SELECT id, file_name, extension, folder_path
	FROM workbooks
	WHERE id = ?
	`, id)

	err := row.Scan(&workbook.ID, &workbook.FileName, &workbook.Extension, &workbook.FolderPath)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("metadata not found for ID: %s", id)
		}
		return nil, fmt.Errorf("failed to retrieve metadata: %v", err)
	}
	return &workbook, nil
}

func (r *WorkbookRepository) DeleteMetadata(id string) error {
	_, err := r.DB.Exec(`
	DELETE FROM workbooks
	WHERE id = ?
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete metadata: %v", err)
	}
	return nil
}

func (r *WorkbookRepository) Close() {
	r.DB.Close()
}

func (r *WorkbookRepository) GetAllWorkbooks() ([]*models.Workbook, error) {
	var workbookList []*models.Workbook

	rows, err := r.DB.Query(`
	SELECT id, file_name, extension, folder_path
	FROM workbooks
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve metadata: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var workbook models.Workbook
		err := rows.Scan(&workbook.ID, &workbook.FileName, &workbook.Extension, &workbook.FolderPath)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve metadata: %v", err)
		}
		workbookList = append(workbookList, &workbook)
	}
	return workbookList, nil
}

// debug function to print all metadata
func (r *WorkbookRepository) PrintWorkbookList() {
	metadataList, err := r.GetAllWorkbooks()
	if err != nil {
		fmt.Printf("failed to retrieve metadata: %v", err)
		return
	}

	for _, metadata := range metadataList {
		fmt.Printf("ID: %s\n", metadata.ID)
		fmt.Printf("File Name: %s\n", metadata.FileName)
		fmt.Printf("Extension: %s\n", metadata.Extension)
		fmt.Printf("Folder Path: %s\n", metadata.FolderPath)
		fmt.Println("----------------------------")
	}
}
