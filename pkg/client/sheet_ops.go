package client

import (
	"fmt"
	"os"

	"github.com/Deathfireofdoom/excel-client-go/pkg/excel"
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
)

func (c *ExcelClient) CreateSheet(workbookID, sheetName string) (*models.Sheet, error) {
	// get metadata of workbook from db
	workbook, err := c.repository.GetMetadata(workbookID)
	if err != nil {
		fmt.Printf("failed to get metadata: %v", err)
		return nil, err
	}

	// check if file exists
	if _, err := os.Stat(workbook.GetFullPath()); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %v", err)
		return nil, err
	}

	// create sheet
	pos, err := excel.CreateSheet(*workbook, sheetName)
	if err != nil {
		fmt.Printf("failed to create sheet: %v", err)
		return nil, err
	}

	// creates sheet object
	sheet, err := models.NewSheet(workbookID, pos, sheetName, "")
	if err != nil {
		fmt.Printf("failed to create sheet: %v", err)
		return nil, err
	}

	// save sheet to db
	err = c.repository.SaveSheet(sheet)
	if err != nil {
		fmt.Printf("failed to save sheet: %v", err)
		return nil, err
	}

	return sheet, nil
}

func (c *ExcelClient) ReadSheet(workbookID, sheetID string) (*models.Sheet, error) {
	// get metadata of workbook from db
	workbook, err := c.repository.GetMetadata(workbookID)
	if err != nil {
		fmt.Printf("failed to get metadata: %v", err)
		return nil, err
	}

	// check if file exists
	if _, err := os.Stat(workbook.GetFullPath()); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %v", err)
		return nil, err
	}

	// get sheet from db
	sheet, err := c.repository.GetSheet(sheetID)
	if err != nil {
		fmt.Printf("failed to get sheet: %v", err)
		return nil, err
	}

	// get sheet from file
	sheet, err = excel.GetSheet(*workbook, sheet.Name, sheet.ID)
	if err != nil {
		fmt.Printf("failed to get sheet: %v", err)
		return nil, err
	}

	return sheet, nil
}

func (c *ExcelClient) DeleteSheet(workbookID, sheetID string) error {
	// get metadata of workbook from db
	workbook, err := c.repository.GetMetadata(workbookID)
	if err != nil {
		fmt.Printf("failed to get metadata: %v", err)
		return err
	}

	// check if file exists
	if _, err := os.Stat(workbook.GetFullPath()); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %v", err)
		return err
	}

	// get sheet from db
	sheet, err := c.repository.GetSheet(sheetID)
	if err != nil {
		fmt.Printf("failed to get sheet: %v", err)
		return err
	}

	// delete sheet from file
	err = excel.DeleteSheet(*workbook, sheet.Name)
	if err != nil {
		fmt.Printf("failed to delete sheet: %v", err)
		return err
	}

	// delete sheet from db
	err = c.repository.DeleteSheet(sheetID)
	if err != nil {
		fmt.Printf("failed to delete sheet: %v", err)
		return err
	}

	return nil
}

func (c *ExcelClient) UpdateSheet(sheet *models.Sheet) (*models.Sheet, error) {
	// get metadata of workbook from db
	workbook, err := c.repository.GetMetadata(sheet.WorkbookID)
	if err != nil {
		fmt.Printf("failed to get metadata: %v", err)
		return nil, err
	}

	// check if file exists
	if _, err := os.Stat(workbook.GetFullPath()); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %v", err)
		return nil, err
	}

	// get sheet from db
	oldSheet, err := c.repository.GetSheet(sheet.ID)
	if err != nil {
		fmt.Printf("failed to get sheet: %v", err)
		return nil, err
	}

	// update sheet in file
	err = excel.RenameSheet(*workbook, oldSheet.Name, sheet.Name)
	if err != nil {
		fmt.Printf("failed to update sheet: %v", err)
		return nil, err
	}

	// update sheet in db
	err = c.repository.SaveSheet(sheet)
	if err != nil {
		fmt.Printf("failed to update sheet in db: %v", err)
		return nil, err
	}

	return sheet, nil
}
