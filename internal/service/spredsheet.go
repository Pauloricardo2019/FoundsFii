package service

import (
	"github.com/tealeg/xlsx"
)

type spreadsheetService struct {
}

func NewSpreadsheetService() *spreadsheetService {
	return &spreadsheetService{}
}

func (spreadsheetService) ReadSpreedsheet() (chan string, error) {

	file, err := xlsx.OpenFile("./spreadsheet/FiiToSearch.xlsx")
	if err != nil {
		return nil, err
	}

	value := make(chan string)

	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {

				value <- cell.Value

			}
		}
	}

	return value, nil
}
