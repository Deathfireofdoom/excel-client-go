package main

import (
	"fmt"

	"github.com/Deathfireofdoom/excel-client-go/pkg/client"
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
)

func main() {
	fmt.Printf("Hello, this is a test to try out the client. Dont mind me.\n")
	excelClient, err := client.NewExcelClient()
	if err != nil {
		fmt.Printf("failed to create client: %v", err)
		return
	}

	// test config
	fileName := "excel_test_file_cell"
	extension := models.Extension("xlsx")
	filePath := "/Users/oskarelvkull/Documents/test-drive/"

	// // create excel file
	workbook, _ := excelClient.CreateWorkbook(filePath, fileName, string(extension), "")
	excelClient.PrintWorkbookList()

	// // create sheet
	sheetName := "test_sheet"
	sheet, err := excelClient.CreateSheet(workbook.ID, sheetName)
	if err != nil {
		fmt.Printf("failed to create sheet: %v", err)
		return
	}
	fmt.Printf("sheet created: %v\n", sheet)

	// // create cell
	cell, err := excelClient.CreateCell(workbook.ID, sheet.ID, 1, "A", "test_value")
	if err != nil {
		fmt.Printf("failed to create cell: %v", err)
		return
	}
	fmt.Printf("cell created: %v\n", cell)

}
