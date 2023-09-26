package excel

import (
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
	excelize "github.com/xuri/excelize/v2"

	"fmt"
)

func CreateSheet(workbook models.Workbook, sheetName string) (int, error) {
	// opens the file
	file, err := excelize.OpenFile(workbook.GetFullPath())
	if err != nil {
		return -1, err
	}
	defer file.Close()

	// creates the sheet
	var sheetNumber int
	if sheetNumber, err = file.NewSheet(sheetName); err != nil {
		return -1, err
	}

	// saves the file
	if err := file.Save(); err != nil {
		fmt.Printf("failed to save file: %v", err)
		return -1, err
	}

	return sheetNumber, nil
}

func DeleteSheet(workbook models.Workbook, sheetName string) error {
	// opens the file
	file, err := excelize.OpenFile(workbook.GetFullPath())
	if err != nil {
		return err
	}
	defer file.Close()

	// removes the sheet
	if err := file.DeleteSheet(sheetName); err != nil {
		return err
	}

	// saves the file
	if err := file.Save(); err != nil {
		fmt.Printf("failed to save file: %v", err)
		return err
	}

	return nil
}

func RenameSheet(workbook models.Workbook, oldSheetName, newSheetName string) error {
	// opens the file
	file, err := excelize.OpenFile(workbook.GetFullPath())
	if err != nil {
		return err
	}
	defer file.Close()

	// renames the sheet
	if err := file.SetSheetName(oldSheetName, newSheetName); err != nil {
		fmt.Printf("failed to rename sheet: %v", err)
		return err
	}

	// saves the file
	if err := file.Save(); err != nil {
		fmt.Printf("failed to save file: %v", err)
		return err
	}

	return nil
}

func GetSheet(workbook models.Workbook, sheetName, sheetID string) (*models.Sheet, error) {
	// opens the file
	file, err := excelize.OpenFile(workbook.GetFullPath())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// gets the sheet
	pos, err := file.GetSheetIndex(sheetName)
	if err != nil {
		return nil, err
	}

	// creates the sheet
	sheet, err := models.NewSheet(workbook.ID, pos, sheetName, sheetID)
	if err != nil {
		return nil, err
	}
	return sheet, nil
}
