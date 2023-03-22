package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"sync"
)

const pattern = "(?P<Liquidez>([0-9]+)||([0-9]+.[0-9]+))\\s+(?P<UltimoRendimento>R\\$.+)\\s+(?P<DividendYield>[0-9]+\\,[0-9]+\\%)\\s+(?P<PatrimonioLiquido>R\\$\\s[0-9]+.[0-9]+\\s\\w+)\\s+(?P<ValorPatrimonial>R\\$\\s[0-9]+\\,[0-9]+)\\s+(?P<RentabNoMes>.[0-9]+\\,[0-9]+\\%)\\s+(?P<PVP>[0-9]+\\,[0-9]+)"

var (
	re *regexp.Regexp
)

func getFii(code string, group *sync.WaitGroup) error {

	defer group.Done()

	url := fmt.Sprintf("https://www.fundsexplorer.com.br/funds/%s", code)

	regex := regexp.MustCompile(pattern)

	re = regex

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	values := doc.Find("div.container").Find("span.indicator-value").Text()
	if err != nil {
		panic(err)
	}

	matches := re.FindStringSubmatch(values)

	groups := re.SubexpNames()
	liquidez := getMatchedValueByIdentifier("Liquidez", matches, groups)
	ultRend := getMatchedValueByIdentifier("UltimoRendimento", matches, groups)
	divYield := getMatchedValueByIdentifier("DividendYield", matches, groups)
	patLiquido := getMatchedValueByIdentifier("PatrimonioLiquido", matches, groups)
	valPatrimonial := getMatchedValueByIdentifier("ValorPatrimonial", matches, groups)
	rentMes := getMatchedValueByIdentifier("RentabNoMes", matches, groups)

	fmt.Printf("Liquidez: %s \nUltimo rendimento: %s \nDividendYeld: %s \nPatrimonio Liquido: %s \nValor Patrimonial: %s \nRentabilidade ao mes: %s \n\n\n", liquidez, ultRend, divYield, patLiquido, valPatrimonial, rentMes)
	return nil
}

func main() {

	wg := &sync.WaitGroup{}

	fiis := []string{
		"mxrf11",
		"aazq11",
		"abcp11",
		"afhi11",
		"kisu11",
		"snff11",
		"snci11",
		"recr11",
		"xpci11",
	}

	maxRoutines := len(fiis)

	wg.Add(maxRoutines)

	for _, fii := range fiis {
		go getFii(fii, wg)
	}

	wg.Wait()

}

func getMatchedValueByIdentifier(id string, matches []string, groups []string) string {
	for _, v := range groups {
		if v == id {
			idx := re.SubexpIndex(v)
			return matches[idx]
		}
	}
	return ""
}
