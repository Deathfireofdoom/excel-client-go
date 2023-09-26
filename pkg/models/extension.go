// Package models contains the data models used in the application.
package models

// Extension represents the type of file extension for a workbook.
type Extension string

// File extension types for Excel workbooks.
const (
	XLSX  Extension = "xlsx"  // Standard Excel workbook
	XLSXM Extension = "xlsm"  // Excel workbook with macros
	XLS   Extension = "xls"   // Excel 97-2003 workbook
)
