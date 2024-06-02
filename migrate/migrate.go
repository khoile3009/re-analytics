package main

import (
	"github.com/khoile3009/re-analytics/initializers"
	"github.com/khoile3009/re-analytics/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Property{})
}
