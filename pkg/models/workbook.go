package models

import (
	"fmt"
	"path/filepath"

	// Third-party libraries
	"github.com/Deathfireofdoom/excel-client-go/pkg/utils"
)

// Workbook represents a spreadsheet file containing multiple sheets.
// It has an ID, file name, extension, and a folder path where it is located.
type Workbook struct {
	ID         string    `json:"id"`
	FileName   string    `json:"file_name"`
	Extension  Extension `json:"extension"`  // File extension (e.g., xls, xlsx)
	FolderPath string    `json:"folder_path"`// Folder path where the workbook is located
	Sheets     []Sheet   `json:"sheets"`     // Sheets within the workbook
}

// GetFullPath constructs and returns the full path of the workbook.
func (e *Workbook) GetFullPath() string {
	fileNameWithExtension := fmt.Sprintf("%s.%s", e.FileName, e.Extension)
	return filepath.Join(e.FolderPath, fileNameWithExtension)
}

// NewWorkbook initializes a new Workbook instance.
// If id is not provided, a new UUID will be generated.
func NewWorkbook(fileName string, extension Extension, folderPath, id string) (*Workbook, error) {
	if id == "" {
		var err error
		id, err = utils.GenerateUUID()
		if err != nil {
			return nil, fmt.Errorf("failed to generate UUID for Workbook: %w", err)
		}
	}

	return &Workbook{
		ID:         id,
		FileName:   fileName,
		Extension:  extension,
		FolderPath: folderPath,
		Sheets:     nil,
	}, nil
}
