package main

import (
	"github.com/khoile3009/re-analytics/initializers"
	"github.com/khoile3009/re-analytics/scrapper"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	scrapper.GetAllRedfinListing(78705)
}
