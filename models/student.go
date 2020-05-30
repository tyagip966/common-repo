package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name           string `json:"name"`
	Age            string `json:"age"`
	Standard       string `json:"standard"`
	SchoolCode     int    `json:"schoolCode"`
	IdentityNumber int    `json:"identityNumber"`
}
