package models

import (
	"path/filepath"

	"github.com/Deathfireofdoom/excel-client-go/pkg/utils"
)

type Workbook struct {
	ID         string    `json:"id"`
	FileName   string    `json:"file_name"`
	Extension  Extension `json:"extension"`
	FolderPath string    `json:"folder_path"`
	Sheets     []Sheet   `json:"sheets"`
}

func (e *Workbook) GetFullPath() string {
	fileNameWithExtension := e.FileName + "." + string(e.Extension)
	fullPath := filepath.Join(e.FolderPath, fileNameWithExtension)
	return fullPath
}

func NewWorkbook(fileName string, extension Extension, folderPath, id string) (*Workbook, error) {
	// generate uuid
	if id == "" {
		var err error
		id, err = utils.GenerateUUID()
		if err != nil {
			return nil, err
		}
	}

	return &Workbook{
		ID:         id,
		FileName:   fileName,
		Extension:  extension,
		FolderPath: folderPath,
	}, nil
}
