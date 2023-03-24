package service

import (
	"WhatsTheBestFii/internal/model"
	"github.com/tealeg/xlsx"
	"golang.org/x/sync/errgroup"
	"log"
	"sync"
	"time"
)

type spreadsheetService struct {
}

func NewSpreadsheetService() *spreadsheetService {
	return &spreadsheetService{}
}

var (
	mu sync.Mutex
)

func (spreadsheetService) ReadSpreedsheet() (chan string, error) {

	file, err := xlsx.OpenFile("./spreadsheet/FiiToSearch.xlsx")
	if err != nil {
		return nil, err
	}

	chanValue := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		for _, sheet := range file.Sheets {
			for _, row := range sheet.Rows {
				for _, cell := range row.Cells {

					if cell.Value == "" {
						close(chanValue)
						return
					}
					chanValue <- cell.Value

				}
			}
		}
	}()

	return chanValue, nil
}

func (spreadsheetService) CreateSpreedsheet(fiis []model.Fii, sheet *xlsx.Sheet) error {

	header := []string{
		"Liquidity",
		"Last Income",
		"Dividend Yield",
		"Net Worth",
		"Book Value",
		"Profitability Per Month",
		"PVP",
	}

	row := sheet.AddRow()
	row.WriteSlice(&header, len(header))

	for _, fii := range fiis {
		g := new(errgroup.Group)

		g.Go(func() error {
			mu.Lock()
			r := sheet.AddRow()
			r.WriteStruct(&fii, -1)
			mu.Unlock()
			return nil
		})

		err := g.Wait()
		if err != nil {
			log.Println("Err on wait group", err.Error())
			return err
		}
	}

	if err := sheet.File.Save("./spreadsheet/FiisFounds.xlsx"); err != nil {
		return err
	}

	return nil
}
