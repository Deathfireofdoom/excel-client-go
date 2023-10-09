package models_test

import (
	"testing"

	"github.com/Deathfireofdoom/excel-client-go/pkg/models" // replace with your actual import path
)

func TestNewCellWithProvidedID(t *testing.T) {
	workbookID := "workbook1"
	sheetID := "sheet1"
	row := 1
	column := "A"
	value := "test"
	id := "uuid1"

	cell, err := models.NewCell(workbookID, sheetID, row, column, value, id)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cell.ID != id {
		t.Errorf("expected id to be %s, got: %s", id, cell.ID)
	}
}

func TestNewCellWithGeneratedID(t *testing.T) {
	workbookID := "workbook1"
	sheetID := "sheet1"
	row := 1
	column := "A"
	value := "test"

	cell, err := models.NewCell(workbookID, sheetID, row, column, value, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cell.ID == "" {
		t.Error("expected id but got nothing")
	}
	// You could also add more assertions here to check the format of the generated ID,
	// or any other relevant properties.
}

func TestGetPosition(t *testing.T) {
	cell := &models.Cell{
		Row:    1,
		Column: "A",
	}
	expectedPosition := "A1"
	if cell.GetPosition() != expectedPosition {
		t.Errorf("expected position to be %s, got: %s", expectedPosition, cell.GetPosition())
	}
}
