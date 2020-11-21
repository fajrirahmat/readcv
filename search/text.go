package search

import (
	"fmt"
	"regexp"
	"strings"
)

var golangKeywords = []string{
	"Golang",
	"Go",
	"Go Programming Language",
	"GO Lang",
	"Go Language",
}

//IsFoundGolang ...
func IsFoundGolang(text string) (bool, error) {
	reg, err := regexp.Compile(generateGolangKeyword(golangKeywords))
	if err != nil {
		return false, err
	}
	return reg.Match([]byte(text)), nil
}

//FoundKeyword ...
func FoundKeyword(text string, keywords []string) (bool, error) {
	reg, err := regexp.Compile(generateGolangKeyword(keywords))
	if err != nil {
		return false, err
	}
	return reg.Match([]byte(text)), nil
}

func generateGolangKeyword(keywords []string) string {
	var temp []string
	for _, golangKeyword := range keywords {
		temp = append(temp, golangKeyword)
		temp = append(temp, strings.ToLower(golangKeyword))
		temp = append(temp, strings.ToUpper(golangKeyword))
	}
	return fmt.Sprintf("(?:%s)", strings.Join(temp, "|"))
}
