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
	fileName := "excel_test_file_2"
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

	// // rename sheet
	newSheetName := "test_sheet_renamed"
	sheet.Name = newSheetName
	_, err = excelClient.UpdateSheet(workbook.ID, sheet)
	if err != nil {
		fmt.Printf("failed to rename sheet: %v", err)
		return
	}
	fmt.Printf("sheet renamed: %v\n", newSheetName)

	// // delete sheet
	err = excelClient.DeleteSheet(workbook.ID, sheet.ID)
	if err != nil {
		fmt.Printf("failed to delete sheet: %v", err)
		return
	}
	fmt.Printf("sheet deleted: %v\n", newSheetName)
}
