package service

import (
	"WhatsTheBestFii/internal/constants"
	"WhatsTheBestFii/internal/model"
	"regexp"
)

type regexService struct {
}

func NewRegexService() *regexService {
	return &regexService{}
}

var (
	re *regexp.Regexp
)

func (regexService) init() {
	regex := regexp.MustCompile(constants.RegexPattern)
	re = regex
}

func (r *regexService) SeparateByGroup(values string) (*model.Fii, error) {
	r.init()

	matches := re.FindStringSubmatch(values)
	groups := re.SubexpNames()
	liquidez := r.getMatchedValueByIdentifier("Liquidez", matches, groups)
	ultRend := r.getMatchedValueByIdentifier("UltimoRendimento", matches, groups)
	divYield := r.getMatchedValueByIdentifier("DividendYield", matches, groups)
	patLiquido := r.getMatchedValueByIdentifier("PatrimonioLiquido", matches, groups)
	valPatrimonial := r.getMatchedValueByIdentifier("ValorPatrimonial", matches, groups)
	rentMes := r.getMatchedValueByIdentifier("RentabNoMes", matches, groups)
	pvp := r.getMatchedValueByIdentifier("PVP", matches, groups)

	return &model.Fii{
		Liquidity:             liquidez,
		LastIncome:            ultRend,
		DividendYield:         divYield,
		NetWorth:              patLiquido,
		BookValue:             valPatrimonial,
		ProfitabilityPerMonth: rentMes,
		PVP:                   pvp,
	}, nil
}

func (regexService) getMatchedValueByIdentifier(id string, matches []string, groups []string) string {
	for _, v := range groups {
		if v == id {
			idx := re.SubexpIndex(v)
			return matches[idx]
		}
	}
	return ""
}
