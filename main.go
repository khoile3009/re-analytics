package main

import (
	"fmt"

	"github.com/khoile3009/re-analytics/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Print("Start application")
}
