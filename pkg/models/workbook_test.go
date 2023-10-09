package models_test

import (
	"path/filepath"
	"testing"

	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
)

func TestNewWorkbookWithProvidedID(t *testing.T) {
	name := "workbook1"
	folderPath := "folder1"
	extension := "xlsx"
	id := "uuid1"

	workbook, err := models.NewWorkbook(name, models.Extension(extension), folderPath, id)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if workbook.ID != id {
		t.Errorf("expected id to be %s, got: %s", id, workbook.ID)
	}
}

func TestNewWorkbookWithoutID(t *testing.T) {
	name := "workbook1"
	folderPath := "folder1"
	extension := "xlsx"

	workbook, err := models.NewWorkbook(name, models.Extension(extension), folderPath, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if workbook.ID == "" {
		t.Error("expected id but got nothing")
	}
}

func TestGetFullPath(t *testing.T) {
	name := "workbook1"
	folderPath := "folder1"
	extension := "xlsx"
	id := "uuid1"

	workbook, err := models.NewWorkbook(name, models.Extension(extension), folderPath, id)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expectedFullPath := filepath.Join(folderPath, name+"."+extension)
	if workbook.GetFullPath() != expectedFullPath {
		t.Errorf("expected full path to be %s, got: %s", expectedFullPath, workbook.GetFullPath())
	}
}
