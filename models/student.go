package models

type Student struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Age            string `json:"age"`
	Standard       string `json:"standard"`
	SchoolCode     int    `json:"schoolCode"`
	IdentityNumber int    `json:"identityNumber"`
}

type StudentInput struct {
	Name           string `json:"name"`
	Age            string `json:"age"`
	Standard       string `json:"standard"`
	SchoolCode     int    `json:"schoolCode"`
	IdentityNumber int    `json:"identityNumber"`
}