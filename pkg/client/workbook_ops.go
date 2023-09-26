package client

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Deathfireofdoom/excel-client-go/pkg/excel"
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
)

// CreateExcel creates an excel file in the specified folder path with the specified file name and extension
func (c *ExcelClient) CreateWorkbook(folderPath, fileName, extension, id string) (*models.Workbook, error) {
	// creates object with metadata
	workbook, err := models.NewWorkbook(fileName, models.Extension(extension), folderPath, id)
	if err != nil {
		fmt.Printf("failed to create metadata-object: %v", err)
		return nil, err
	}

	// creates workbook
	workbook, err = excel.CreateWorkbook(workbook)
	if err != nil {
		fmt.Printf("failed to create workbook: %v", err)
		return nil, err
	}

	// save metadata to db
	err = c.repository.SaveMetadata(workbook)
	if err != nil {
		fmt.Printf("failed to save metadata: %v", err)
		return nil, err
	}

	return workbook, nil
}

func (c *ExcelClient) ReadWorkbook(id string) (*models.Workbook, error) {
	// get metadata from db
	workbook, err := c.repository.GetWorkbook(id)
	if err != nil {
		fmt.Printf("failed to get metadata: %v", err)
		return nil, err
	}

	// check if file exists
	if _, err := os.Stat(workbook.GetFullPath()); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %v", err)
		return nil, err
	}

	return workbook, nil
}

func (c *ExcelClient) DeleteWorkbook(id string) error {
	// get metadata from db
	workbook, err := c.repository.GetMetadata(id)
	if err != nil {
		fmt.Printf("failed to get metadata: %v", err)
		return err
	}

	// check if file exists
	if _, err := os.Stat(workbook.GetFullPath()); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %v", err)
		return nil // todo check if this is correct
	}

	// delete file
	err = os.Remove(workbook.GetFullPath())
	if err != nil {
		fmt.Printf("failed to delete file: %v", err)
		return err
	}

	// delete metadata from db
	err = c.repository.DeleteMetadata(id)
	if err != nil {
		fmt.Printf("failed to delete metadata: %v", err)
		return err
	}

	return nil
}

func (c *ExcelClient) UpdateWorkbook(workbook *models.Workbook) (*models.Workbook, error) {
	// get old metadata from db
	oldWorkbook, err := c.repository.GetMetadata(workbook.ID)
	if err != nil {
		fmt.Printf("failed to get metadata: %v\n", err)
		return nil, err
	}

	// check if file exists
	if _, err := os.Stat(oldWorkbook.GetFullPath()); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %v\n", err)
		fmt.Printf("creating new file instead\n")
		workbook, err := c.CreateWorkbook(workbook.FolderPath, workbook.FileName, string(workbook.Extension), workbook.ID)
		if err != nil {
			fmt.Printf("failed to create new file: %v\n", err)
			return nil, err
		}
		return workbook, err
	}

	// create the folders if they does not exists
	err = os.MkdirAll(filepath.Dir(workbook.GetFullPath()), os.ModePerm)
	if err != nil {
		fmt.Printf("failed to create the folder structure: %v", err)
		return nil, err
	}

	// rename file
	err = os.Rename(oldWorkbook.GetFullPath(), workbook.GetFullPath())
	if err != nil {
		fmt.Printf("failed to rename file: %v\n", err)
		return nil, err
	}

	// update metadata in db
	err = c.repository.SaveMetadata(workbook)
	if err != nil {
		fmt.Printf("failed to update metadata: %v\n", err)
		return nil, err
	}

	return workbook, nil
}
