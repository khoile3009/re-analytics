package scrapper

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// This is scrapper for realtor.com

const BaseRealtorUrl = "www.realtor.com"

func GetRedfinUrl(zipCode int) string {
	return fmt.Sprintf("https://www.realtor.com/realestateandhomes-search/%v", zipCode)
}

func GetAllRedfinListing(zipCode int) {
	collector := colly.NewCollector(
	// // Restrict crawling to specific domains
	// colly.AllowedDomains("www.realtor.com"),
	// // Allow visiting the same page multiple times
	// colly.AllowURLRevisit(),
	// // Allow crawling to be done in parallel / async
	// colly.Async(true),
	)
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})
	// Get list of properties
	collector.OnHTML("section", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("class"), "PropertiesList") {
			e.ForEach("div", func(_ int, childDiv *colly.HTMLElement) {
				if strings.HasPrefix(childDiv.Attr("id"), "property_id") {
					contentDiv := childDiv.DOM.Find("div.card-content")
					if contentDiv.Length() == 0 {
						log.Fatal("Couldn't find content div")
					}
					// Get message (condo for sale...)
					messageDiv := contentDiv.Find("div.message")
					if messageDiv.Length() == 0 {
						log.Fatal("Couldn't find message div")
					}
					message := strings.TrimSpace(messageDiv.Text())
					fmt.Println(message)

					// Get price
					priceDiv := contentDiv.Find("div.card-price")
					if priceDiv.Length() == 0 {
						log.Fatal("Couldn't find price div")
					}
					price := ConvertPriceTextToFloat(priceDiv.Text())
					fmt.Println(price)

					// Get bed, bath and size information
					metaUl := contentDiv.Find("ul.card-meta")
					var numberOfBeds float32
					var numberOfBaths float32
					var sizeSquareFeet float32
					var lotSizeSquareFeet float32
					propertyLis := metaUl.Find("li")
					propertyLis.Each(
						func(_ int, propertyLi *goquery.Selection) {
							testId, exists := propertyLi.Attr("data-testid")
							if !exists {
								log.Fatal("data-testid does not exist")
							}
							if testId == "property-meta-beds" {
								span := propertyLi.Find(`span[data-testid="meta-value"]`)
								if span.Length() == 0 {
									log.Fatal("Couldn't find span")
								}
								numberOfBeds = ConvertStringToFloat(span.Text())

							}
							if testId == "property-meta-baths" {
								span := propertyLi.Find(`span[data-testid="meta-value"]`)
								if span.Length() == 0 {
									log.Fatal("Couldn't find span")
								}
								numberOfBaths = ConvertStringToFloat(span.Text())
							}
							if testId == "property-meta-sqft" {
								span := propertyLi.Find(`span[data-testid="meta-value"]`)
								if span.Length() == 0 {
									log.Fatal("Couldn't find span")
								}
								sizeSquareFeet = ConvertStringToFloat(span.Text())
							}
							if testId == "property-meta-lot-size" {
								span := propertyLi.Find(`span[data-testid="meta-value"]`)
								if span.Length() == 0 {
									log.Fatal("Couldn't find span")
								}
								lotSizeSquareFeet = ConvertStringToFloat(span.Text())

							}
						})
					fmt.Println(numberOfBeds)
					fmt.Println(numberOfBaths)
					fmt.Println(sizeSquareFeet)
					fmt.Println(lotSizeSquareFeet)
					log.Fatal("break")
				}
			})
		}

	})
	url := GetRedfinUrl(zipCode)
	fmt.Print(url)
	collector.Visit(url)
}
