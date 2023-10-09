package models_test

import (
	"testing"

	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
)

func TestNewSheetWithoutID(t *testing.T) {
	workbookID := "workbook1"
	pos := 1
	name := "sheet1"
	id := ""

	sheet, err := models.NewSheet(workbookID, pos, name, id)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if sheet.ID == "" {
		t.Error("expected id but got nothing")
	}
}

func TestNewSheetWithProvidedID(t *testing.T) {
	workbookID := "workbook1"
	pos := 1
	name := "sheet1"
	id := "uuid1"

	sheet, err := models.NewSheet(workbookID, pos, name, id)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if sheet.ID != id {
		t.Errorf("expected id to be %s, got: %s", id, sheet.ID)
	}
}
