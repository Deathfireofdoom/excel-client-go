package client

import "github.com/Deathfireofdoom/excel-client-go/pkg/excel"

func (c *ExcelClient) GetExtensions() []string {
	return excel.GetExtensions()
}

func (c *ExcelClient) PrintWorkbookList() {
	c.repository.PrintWorkbookList()
}
