package main

import (
	"WhatsTheBestFii/internal/facade"
	"WhatsTheBestFii/internal/provider"
	"WhatsTheBestFii/internal/service"
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
    //setup 
	spreadsheetService := service.NewSpreadsheetService()
	regexService := service.NewRegexService()

	foundExplorerProvider := provider.NewFundsExplorerProvider()

	processFiiFacade := facade.NewProcessFiiFacade(
		spreadsheetService,
		foundExplorerProvider,
		regexService,
	)

	file := xlsx.NewFile()

	fileFii, err := file.AddSheet("Fii")
	if err != nil {

	}

	if err := processFiiFacade.Process(fileFii); err != nil {
		fmt.Println(err.Error())
	}

}
