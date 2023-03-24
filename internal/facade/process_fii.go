package facade

import (
	"WhatsTheBestFii/internal/model"
	"fmt"
	"github.com/tealeg/xlsx"
)

type processFiiFacade struct {
	spreadsheetService     spreadsheetService
	providerFoundsExplorer providerFoundsExplorer
	regexService           regexService
}

func NewProcessFiiFacade(
	spreadsheetService spreadsheetService,
	providerFoundsExplorer providerFoundsExplorer,
	regexService regexService,
) *processFiiFacade {
	return &processFiiFacade{
		spreadsheetService:     spreadsheetService,
		providerFoundsExplorer: providerFoundsExplorer,
		regexService:           regexService,
	}
}

func (p *processFiiFacade) Process(fileFii *xlsx.Sheet) error {

	valueChan, err := p.spreadsheetService.ReadSpreedsheet()
	if err != nil {
		return err
	}

	fiis := make([]model.Fii, 0)

	for value := range valueChan {

		infos, err := p.providerFoundsExplorer.GetInfos(value)
		if err != nil {
			return err
		}
		fmt.Println("Fii: ", value)
		fii, err := p.regexService.SeparateByGroup(*infos)
		if err != nil {
			return err
		}

		fiis = append(fiis, *fii)
	}

	return p.spreadsheetService.CreateSpreedsheet(fiis, fileFii)

}
