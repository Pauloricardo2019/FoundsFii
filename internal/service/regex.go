package service

import (
	"WhatsTheBestFii/internal/constants"
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

func (regexService) SeparateByGroup(values string) {

}
