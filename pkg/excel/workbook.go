package excel

import (
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
	excelize "github.com/xuri/excelize/v2"

	"fmt"
	"os"
	"path/filepath"
)

func CreateWorkbook(workbook *models.Workbook) (*models.Workbook, error) {
	// create the folders if they does not exists
	err := os.MkdirAll(filepath.Dir(workbook.GetFullPath()), os.ModePerm)
	if err != nil {
		fmt.Printf("failed to create the folder structure: %v", err)
		return nil, err
	}

	// check if file exists
	if _, err := os.Stat(workbook.GetFullPath()); err == nil {
		fmt.Printf("File already exists: %s\n", workbook.GetFullPath())
		return nil, err
	}

	// create the file
	file := excelize.NewFile()
	defer file.Close()

	// create the sheet
	sheetName := "Sheet1"
	file.NewSheet(sheetName)

	// save the file
	if err := file.SaveAs(workbook.GetFullPath()); err != nil {
		return nil, err
	}

	return workbook, nil
}
