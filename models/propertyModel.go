package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Url               string
	Message           *string
	AddressRaw        string
	Address           *string
	City              *string
	State             *string
	ZipCode           *string
	Type              *string
	NumberOfBed       *float32
	NumberOfBath      *float32
	SizeSquareFeet    *float32
	LotSizeSquareFeet *float32
	Price             *float32
}
