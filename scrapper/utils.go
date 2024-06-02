package scrapper

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ConvertPriceTextToFloat(priceText string) float32 {
	priceText = strings.ReplaceAll(priceText, "$", "")
	return ConvertStringToFloat(priceText)
}

func ConvertStringToFloat(text string) float32 {
	text = strings.TrimSpace(text)
	var number float64
	var err error
	text = strings.ReplaceAll(text, ",", "")
	if number, err = strconv.ParseFloat(text, 32); err != nil {
		log.Fatal("Can't convert string to int, price text is " + text)
	}
	return float32(number)
}

func ExtractSizeFromText(text string) int {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindStringSubmatch(text)
	if len(matches) >= 2 {
		// Convert the matched string to an integer
		if intValue, err := strconv.ParseFloat(matches[1], 32); err == nil {
			fmt.Println("Extracted value:", intValue)
		} else {
			fmt.Println("Error converting to integer:", err)
		}
	} else {
		fmt.Println("No numeric value found in the string.")
	}
	return 0
}
