package client

import (
	"fmt"

	"github.com/Deathfireofdoom/excel-client-go/pkg/db"
)

type ExcelClient struct {
	repository *db.WorkbookRepository
}

func NewExcelClient() (*ExcelClient, error) {
	repository, err := db.NewWorkbookRepository()

	if err != nil {
		fmt.Printf("failed to create repository: %v", err)
		return nil, err
	}
	return &ExcelClient{repository: repository}, nil
}
