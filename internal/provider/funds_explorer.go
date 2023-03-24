package provider

import (
	"WhatsTheBestFii/internal/constants"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type fundsExplorerProvider struct {
}

func NewFundsExplorerProvider() *fundsExplorerProvider {
	return &fundsExplorerProvider{}
}

func (fundsExplorerProvider) GetInfos(code string) (*string, error) {

	urlPath := fmt.Sprintf("%s%s", constants.URL, code)

	response, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	values := doc.Find("div.container").Find("span.indicator-value").Text()
	if err != nil {
		panic(err)
	}

	return &values, nil
}
