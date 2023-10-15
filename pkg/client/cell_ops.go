package client

import (
	"fmt"
	"os"

	"github.com/Deathfireofdoom/excel-client-go/pkg/excel"
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
)

func (c *ExcelClient) CreateCell(workbookID, sheetID string, row int, column string, value interface{}) (*models.Cell, error) {
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

	// create cell object
	cell, err := models.NewCell(workbookID, sheetID, row, column, value, "")
	if err != nil {
		fmt.Printf("failed to create cell: %v", err)
		return nil, err
	}

	// create cell in file
	err = excel.CreateCell(workbook, sheet, cell)
	if err != nil {
		fmt.Printf("failed to create cell: %v", err)
		return nil, err
	}

	// save cell to db
	err = c.repository.SaveCell(cell, workbook.ID, sheet.ID)
	if err != nil {
		fmt.Printf("failed to save cell: %v", err)
		return nil, err
	}

	return cell, nil
}

func (c *ExcelClient) ReadCell(workbookID, sheetID, cellID string) (*models.Cell, error) {
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

	// get cell from db
	cell, err := c.repository.GetCell(cellID)
	if err != nil {
		fmt.Printf("failed to get cell: %v", err)
		return nil, err
	}

	// get cell from file
	cell.Value, err = excel.ReadCell(workbook, sheet, cell)
	if err != nil {
		fmt.Printf("failed to get cell: %v", err)
		return nil, err
	}

	return cell, nil
}

func (c *ExcelClient) DeleteCell(workbookID, sheetID, cellID string) error {
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

	// get cell from db
	cell, err := c.repository.GetCell(cellID)
	if err != nil {
		fmt.Printf("failed to get cell: %v", err)
		return err
	}

	// delete cell from file
	err = excel.DeleteCell(workbook, sheet, cell)
	if err != nil {
		fmt.Printf("failed to delete cell: %v", err)
		return err
	}

	// delete cell from db
	err = c.repository.DeleteCell(cellID)
	if err != nil {
		fmt.Printf("failed to delete cell: %v", err)
		return err
	}

	return nil
}

func (c *ExcelClient) UpdateCell(cell *models.Cell) (*models.Cell, error) {
	// get metadata of workbook from db
	workbook, err := c.repository.GetMetadata(cell.WorkbookID)
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
	sheet, err := c.repository.GetSheet(cell.SheetID)
	if err != nil {
		fmt.Printf("failed to get sheet: %v", err)
		return nil, err
	}

	// update cell in file
	err = excel.UpdateCell(workbook, sheet, cell)
	if err != nil {
		fmt.Printf("failed to update cell: %v", err)
		return nil, err
	}

	// update cell in db
	err = c.repository.SaveCell(cell, workbook.ID, sheet.ID)
	if err != nil {
		fmt.Printf("failed to update cell in db: %v", err)
		return nil, err
	}

	return cell, nil
}
