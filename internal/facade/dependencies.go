package facade

import (
	"WhatsTheBestFii/internal/model"
	"github.com/tealeg/xlsx"
)

type spreadsheetService interface {
	CreateSpreedsheet(fiis []model.Fii, sheet *xlsx.Sheet) error
	ReadSpreedsheet() (chan string, error)
}

type providerFoundsExplorer interface {
	GetInfos(code string) (*string, error)
}

type regexService interface {
	SeparateByGroup(values string) (*model.Fii, error)
}
