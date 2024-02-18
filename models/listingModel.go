package models

import "gorm.io/gorm"

type Listing struct {
	gorm.Model
	Name string
	Url  string
}
