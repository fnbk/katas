package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ermittelZeilen(inputStr string) []string {
	if len(inputStr) == 0 {
		return []string{}
	}
	return strings.Split(inputStr, "\n")
}

func ermittleAnzahl(zeilen []string) int {
	return len(zeilen)
}

func berecheneZeilen(inputStr string) int {
	zeilen := ermittelZeilen(inputStr)
	codeZeilen := selektiereCodeZeilen(zeilen)
	return ermittleAnzahl(codeZeilen)
}

func selektiereCodeZeilen(zeilen []string) []string {
	codeZeilen := []string{}
	for i := range zeilen {
		if enthaeltCode(zeilen[i]) {
			codeZeilen = append(codeZeilen, (zeilen[i]))
		}
	}
	return codeZeilen
}

func enthaeltCode(zeile string) bool {
	if len(zeile) > 1 {
		if zeile[0] == '/' && zeile[1] == '/' {
			return false
		}
	}

	strippedZeile := zeile
	strippedZeile = strings.Replace(strippedZeile, " ", "", -1)
	strippedZeile = strings.Replace(strippedZeile, "\t", "", -1)

	if len(strippedZeile) > 0 {
		return true
	}
	return false
}

func leseInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	inputStr := ""
	for scanner.Scan() {
		inputStr = fmt.Sprintf("%s\n%s", inputStr, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	return inputStr
}

func main() {
	inputStr := leseInput()
	count := berecheneZeilen(inputStr)
	fmt.Println(count)
}
