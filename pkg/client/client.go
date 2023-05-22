package client

import (
	"deathfireofdoom/excel-go-client/pkg/db"
	"deathfireofdoom/excel-go-client/pkg/excel"
	"deathfireofdoom/excel-go-client/pkg/models"
	"fmt"
	"os"
	"path/filepath"
)

type excelClient struct {
	repository *db.WorkbookRepository
}

func NewExcelClient() *excelClient {
	repository, err := db.NewWorkbookRepository()
	if err != nil {
		fmt.Printf("failed to create repository: %v", err)
		panic(err)
	}
	return &excelClient{repository: repository}
}

// CreateExcel creates an excel file in the specified folder path with the specified file name and extension
func (c *excelClient) CreateExcel(folderPath, fileName, extension, id string) (*models.Workbook, error) {
	// creates object with metadata
	workbook, err := models.NewWorkbook(fileName, models.Extension(extension), folderPath, id)
	if err != nil {
		fmt.Printf("failed to create metadata-object: %v", err)
		return nil, err
	}

	// create the folders if they does not exists
	err = os.MkdirAll(filepath.Dir(workbook.GetFullPath()), os.ModePerm)
	if err != nil {
		fmt.Printf("failed to create the folder structure: %v", err)
		return nil, err
	}

	// check if file exists
	if _, err := os.Stat(workbook.GetFullPath()); err == nil {
		fmt.Printf("File already exists: %s\n", workbook.GetFullPath())
		return nil, err
	}

	// create file
	file, err := os.Create(workbook.GetFullPath())
	if err != nil {
		fmt.Printf("failed to create the file: %v", err)
		return nil, err
	}
	defer file.Close()

	// save metadata to db
	err = c.repository.SaveMetadata(workbook)
	if err != nil {
		fmt.Printf("failed to save metadata: %v", err)
		return nil, err
	}

	return workbook, nil
}

func (c *excelClient) ReadExcel(id string) (*models.Workbook, error) {
	// get metadata from db
	workbook, err := c.repository.GetMetadata(id)
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

func (c *excelClient) DeleteExcel(id string) error {
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

func (c *excelClient) UpdateExcel(workbook *models.Workbook) (*models.Workbook, error) {
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
		workbook, err := c.CreateExcel(workbook.FolderPath, workbook.FileName, string(workbook.Extension), workbook.ID)
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

func (c *excelClient) GetExtensions() []string {
	return excel.GetExtensions()
}

// debug function to print all metadata
func (c *excelClient) PrintWorkbookList() {
	c.repository.PrintWorkbookList()
}
