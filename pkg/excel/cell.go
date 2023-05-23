package excel

import (
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
	excelize "github.com/xuri/excelize/v2"

	"fmt"
)

func UpdateCell(workbook *models.Workbook, sheet *models.Sheet, cell *models.Cell) error {
	// opens the file
	file, err := excelize.OpenFile(workbook.GetFullPath())
	if err != nil {
		return err
	}
	defer file.Close()

	// set the value
	if err := file.SetCellValue(sheet.Name, cell.GetPosition(), cell.Value); err != nil {
		return err
	}

	// saves the file
	if err := file.Save(); err != nil {
		fmt.Printf("failed to save file: %v", err)
		return err
	}

	return nil
}

func ReadCell(workbook *models.Workbook, sheet *models.Sheet, cell *models.Cell) (interface{}, error) {
	// opens the file
	file, err := excelize.OpenFile(workbook.GetFullPath())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// get the value
	value, err := file.GetCellValue(sheet.Name, cell.GetPosition())
	if err != nil {
		return nil, err
	}

	return value, nil
}

// these are just wrappers for UpdateCell to make it align with the other functions
func DeleteCell(workbook *models.Workbook, sheet *models.Sheet, cell *models.Cell) error {
	cell.Value = nil
	return UpdateCell(workbook, sheet, cell)
}

func CreateCell(workbook *models.Workbook, sheet *models.Sheet, cell *models.Cell) error {
	return UpdateCell(workbook, sheet, cell)
}
